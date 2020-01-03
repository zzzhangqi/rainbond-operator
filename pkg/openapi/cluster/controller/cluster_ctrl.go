package controller

import (
	"net/http"
	"strings"

	"github.com/GLYASAI/rainbond-operator/pkg/openapi/cluster"
	"github.com/GLYASAI/rainbond-operator/pkg/openapi/model"
	"github.com/gin-gonic/gin"
)

// ClusterController k8s controller
type ClusterController struct {
	clusterCase cluster.IClusterCase
}

// NewClusterController creates a new k8s controller
func NewClusterController(g *gin.Engine, clusterCase cluster.IClusterCase) {
	u := &ClusterController{clusterCase: clusterCase}

	clusterEngine := g.Group("/cluster")

	// config
	configEngine := clusterEngine.Group("/configs")
	configEngine.GET("/", u.Configs)
	configEngine.PUT("/", u.UpdateConfig)

	// install
	installEngine := clusterEngine.Group("/install")
	installEngine.POST("/", u.Install)
	installEngine.GET("/status", u.InstallStatus)

	// componse
	componseEngine := clusterEngine.Group("/componses")
	componseEngine.GET("/", u.Componses)
	componseEngine.GET("/:name", u.SingleComponse)
}

// Configs get cluster config info
func (cc *ClusterController) Configs(c *gin.Context) {
	configs, err := cc.clusterCase.GlobalConfigs().GlobalConfigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "msg": "success", "data": configs})
}

// UpdateConfig update cluster config info
func (cc *ClusterController) UpdateConfig(c *gin.Context) {
	var req *model.GlobalConfigs
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := cc.clusterCase.GlobalConfigs().UpdateGlobalConfig(req); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "msg": "success"})
}

// Install install
func (cc *ClusterController) Install(c *gin.Context) {
	if err := cc.clusterCase.Install().Install(); err != nil { // TODO fanyangyang can't download rainbond file filter and return download URL
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "msg": "success"})
}

// InstallStatus install status
func (cc *ClusterController) InstallStatus(c *gin.Context) {
	status, err := cc.clusterCase.Install().InstallStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "msg": "success", "data": map[string]string{"status": status}})
}

// Componses compnses status
func (cc *ClusterController) Componses(c *gin.Context) {
	componseInfos, err := cc.clusterCase.Componses().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "msg": "success", "data": componseInfos})
}

// SingleComponse single componse
func (cc *ClusterController) SingleComponse(c *gin.Context) {
	name := c.Param("name")
	name = strings.TrimSpace(name)
	if name == "" {
		cc.Componses(c) // TODO fanyangyang need for test
		return
	}
	componseInfos, err := cc.clusterCase.Componses().Get(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "msg": "success", "data": componseInfos})
}
