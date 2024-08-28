package service

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/baoer/im_sys/models"
	"github.com/baoer/im_sys/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PrivateMessage struct {
	NickName string `json:"nickname"`
	Uid      string `json:"uid"`
	Msg      string `json:"msg"`
	ToUserID string `json:"touserid"`
}

type Message struct {
	NickName     string `json:"nickname"`
	Uid          string `json:"uid"`
	Msg          string `json:"msg"`
	RoomIdentity string `json:"room_identity"`
}

var upgrader = websocket.Upgrader{}
var prwsc = make(map[string]*websocket.Conn, 0)
var pbwsc = make(map[string]*websocket.Conn, 0)
var prmu sync.Mutex // 用于保护wsc的互斥锁
var pbmu sync.Mutex // 用于保护wsc的互斥锁

func WebSocketsendPrivateMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wsn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer wsn.Close()
		myclaim := ctx.MustGet("user_claims").(*util.Myclaims)
		prmu.Lock()
		prwsc[myclaim.UserId] = wsn
		prmu.Unlock()
		defer func() {
			prmu.Lock()
			delete(prwsc, myclaim.UserId)
			prmu.Unlock()
		}()

		pmessage := new(PrivateMessage)
		for {
			pmessage.Uid = myclaim.UserId
			err := wsn.ReadJSON(pmessage)
			if err != nil {
				log.Println(err.Error())
				break
			}
			msg := models.PrivateMessageBasic{
				Identity:            primitive.NewObjectID(),
				UserIdentity:        pmessage.Uid,
				ReceiveUserIdentity: pmessage.ToUserID,
				Data:                pmessage.Msg,
				CreatedAt:           time.Now().Unix(),
			}
			err = models.InsertOnePrivateMsg(msg)
			if err != nil {
				log.Println(err.Error())
				break
			}
			prmu.Lock()
			if v, ok := prwsc[pmessage.ToUserID]; ok {
				err = v.WriteJSON(pmessage)
				if err != nil {
					log.Println(err.Error())
					break
				}
				err = wsn.WriteJSON(pmessage)
				if err != nil {
					log.Println(err.Error())
					break
				}
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "用户未在线",
				})
			}
			prmu.Unlock()

		}
	}
}

func WebSocketsendChannelMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wsn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer wsn.Close()

		myclaim := ctx.MustGet("user_claims").(*util.Myclaims)
		pbmu.Lock()
		pbwsc[myclaim.UserId] = wsn
		pbmu.Unlock()
		defer func() {
			pbmu.Lock()
			delete(pbwsc, myclaim.UserId)
			pbmu.Unlock()
		}()

		message := new(Message)
		for {
			message.Uid = myclaim.UserId
			err := wsn.ReadJSON(message)
			if err != nil {
				log.Println(err.Error())
				break
			}
			//查询同一房间的用户
			msg := models.PublicMessageBasic{
				Identity:      primitive.NewObjectID(),
				UserIdentity:  message.Uid,
				Room_identity: message.RoomIdentity,
				Data:          message.Msg,
				CreatedAt:     time.Now().Unix(),
			}
			err = models.InsertOnePublicMsg(msg)
			if err != nil {
				log.Println(err.Error())
				break
			}
			users, err := models.GetUsersByRoomIdentity(message.RoomIdentity)
			if err != nil {
				log.Println(err.Error())
				break
			}
			pbmu.Lock()
			for _, uid := range users {
				if v, ok := pbwsc[uid]; ok {
					err = v.WriteJSON(message)
					if err != nil {
						log.Println(err.Error())
						break
					}
				}
			}
			pbmu.Unlock()
		}
	}
}
