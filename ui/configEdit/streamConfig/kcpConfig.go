package streamConfig

import (
	"github.com/therecipe/qt/widgets"
)

type KcpConfig struct {
	*widgets.QWidget

	mtuLineEdit              *widgets.QLineEdit
	ttiLineEdit              *widgets.QLineEdit
	uplinkCapacityLineEdit   *widgets.QLineEdit
	downlinkCapacityLineEdit *widgets.QLineEdit
	readBufferSizeLineEdit   *widgets.QLineEdit
	writeBufferSizeLineEdit  *widgets.QLineEdit

	congestionCheckBox *widgets.QCheckBox

	typeComboBox *widgets.QComboBox
}

func NewKcpConfig(parent widgets.QWidget_ITF) *KcpConfig {
	widget := widgets.NewQWidget(parent, 0)

	kcpConfig := &KcpConfig{QWidget: widget}
	kcpConfig.init()

	return kcpConfig
}

func (ptr *KcpConfig) init() {
	formLayout := widgets.NewQFormLayout(ptr)
	formLayout.SetContentsMargins(0, 0, 0, 0)

	ptr.mtuLineEdit = widgets.NewQLineEdit(ptr)
	ptr.ttiLineEdit = widgets.NewQLineEdit(ptr)
	ptr.uplinkCapacityLineEdit = widgets.NewQLineEdit(ptr)
	ptr.downlinkCapacityLineEdit = widgets.NewQLineEdit(ptr)
	ptr.readBufferSizeLineEdit = widgets.NewQLineEdit(ptr)
	ptr.writeBufferSizeLineEdit = widgets.NewQLineEdit(ptr)

	ptr.mtuLineEdit.SetPlaceholderText("1350")
	ptr.ttiLineEdit.SetPlaceholderText("50")
	ptr.uplinkCapacityLineEdit.SetPlaceholderText("5")
	ptr.downlinkCapacityLineEdit.SetPlaceholderText("20")
	ptr.readBufferSizeLineEdit.SetPlaceholderText("2")
	ptr.writeBufferSizeLineEdit.SetPlaceholderText("2")

	ptr.congestionCheckBox = widgets.NewQCheckBox2("启用", ptr)

	ptr.typeComboBox = widgets.NewQComboBox(ptr)
	ptr.typeComboBox.AddItems([]string{"none", "srtp", "utp", "wechat-video", "dtls", "wireguard"})

	formLayout.AddRow3("最大传输单元(mtu)", ptr.mtuLineEdit)
	formLayout.AddRow3("传输时间间隔(tti)", ptr.ttiLineEdit)
	formLayout.AddRow3("上行链路容量(uplinkCapacity)", ptr.uplinkCapacityLineEdit)
	formLayout.AddRow3("下行链路容量(downlinkCapacity)", ptr.downlinkCapacityLineEdit)
	formLayout.AddRow3("启用拥塞控制(congestion)", ptr.congestionCheckBox)
	formLayout.AddRow3("readBufferSize", ptr.readBufferSizeLineEdit)
	formLayout.AddRow3("writeBufferSize", ptr.writeBufferSizeLineEdit)
	formLayout.AddRow3("伪装类型(header type)", ptr.typeComboBox)

	ptr.SetLayout(formLayout)
}