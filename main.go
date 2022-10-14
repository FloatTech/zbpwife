// Package wife 抽老婆
package wife

import (
	"encoding/json"
	"strconv"

	fcext "github.com/FloatTech/floatbox/ctxext"
	"github.com/FloatTech/floatbox/file"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	"github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type cardInfo struct {
	Name  string   `json:"name"`
	URL   string   `json:"url"`
	Lines []string `json:"lines"`
}
type cardSet = map[string]cardInfo

var cardMap = make(cardSet, 50)
var datapath string

func init() {
	engine := control.Register("wife", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "抽老婆\n" +
			"- 抽老婆",
		PublicDataFolder: "Wives",
	}).ApplySingle(ctxext.DefaultSingle)
	engine.OnFullMatchGroup([]string{"抽老婆"}, fcext.DoOnceOnSuccess(
		func(ctx *zero.Ctx) bool {
			data, err := engine.GetLazyData("wife.json", true)
			if err != nil {
				ctx.SendChain(message.Text("ERROR:", err))
				return false
			}
			err = json.Unmarshal(data, &cardMap)
			if err != nil {
				ctx.SendChain(message.Text("ERROR:", err))
				return false
			}
			datapath = file.BOTPATH + "/" + engine.DataFolder()
			logrus.Infof("[wife]读取%d个老婆", len(cardMap))
			return true
		},
	)).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			i := fcext.RandSenderPerDayN(ctx.Event.UserID, len(cardMap))
			card := cardMap[(strconv.Itoa(i))]

			if id := ctx.SendChain(
				message.At(ctx.Event.UserID),
				message.Text("今天的二次元老婆是~【", card.Name, "】哒"),
				message.Image("file:///"+datapath+card.URL),
				// message.Text("\n【", card.Lines[rand.Intn(len(card.Lines))], "】"),
			); id.ID() == 0 {
				ctx.SendChain(
					message.At(ctx.Event.UserID),
					message.Text("今天的二次元老婆是~【", card.Name, "】哒\n【图片发送失败，请联系维护者~】"))
			}
		})
}
