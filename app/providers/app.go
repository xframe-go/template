package providers

import (
	"github.com/dromara/carbon/v2"
	"github.com/xframe-go/template/app/http/dao"
	"github.com/xframe-go/x/contracts"
	"time"
)
	
type AppServiceProvider struct{}

func (a AppServiceProvider) Register(app contracts.Application) {
	tx, err := app.DB().Connection().DB()
	if err != nil {
		app.Log().Fatal(err.Error())
	}
	if err := tx.Ping(); err != nil {
		return
	}

	carbon.SetDefault(carbon.Default{
		Layout:   time.DateTime,
		Timezone: time.Local.String(),
	})

	dao.SetDefault(app.DB().Connection())
}
