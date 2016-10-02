import QtQuick 2.1
import QtWebEngine 1.2
import QtQuick.Layouts 1.0
import QtQuick.Controls 1.2

ApplicationWindow {
    id: root
    width: 640
    height: 480
    visible: true
	color: "#f8f8f8"
    property int margin: 11
    title: webEngineView && webEngineView.title

    Action {
        id: focus
        shortcut: "Ctrl+L"
        onTriggered: {
            uriField.forceActiveFocus();
            uriField.selectAll();
        }
    }
    Action {
        shortcut: "Ctrl+R"
        onTriggered: {
			webEngineView.reload()
        }
    }
    Action {
        shortcut: "Ctrl+W"
        onTriggered: {
			root.close()
        }
    }

    ColumnLayout {
        id: mainLayout
        anchors.fill: parent
        anchors.margins: margin

		RowLayout {
			id: rowLayout
			anchors.fill: parent

			Rectangle {
				id: backButton
				objectName: "backButton"
				width: 24
				height: 24
				color: backButtonMouseArea.containsMouse ? "#eaeaea" : "transparent"
				signal clicked

				Image {
					anchors.fill: parent
					source: "qrc:///resources/images/button_back.png"
			 	}

				MouseArea {
					id: backButtonMouseArea
					hoverEnabled: true
					anchors.fill: parent
					onClicked: { parent.clicked();}
					onEntered: { parent.color = "#eaeaea" }
					onExited: { parent.color = "transparent" }
					onPressed: { parent.color = "#c7c7c7" }
					onReleased: {
						if (containsMouse) {
							parent.color = "#eaeaea"
						} else {
							parent.color = "transparent"
						}
					}
				}
			}

			Rectangle {
				id: forwardButton
				objectName: "forwardButton"
				width: 24
				height: 24
				color: forwardButtonMouseArea.containsMouse ? "#eaeaea" : "transparent"
				signal clicked

				Image {
					anchors.fill: parent
					source: "qrc:///resources/images/button_forward.png"
			 	}

				MouseArea {
					id: forwardButtonMouseArea
					hoverEnabled: true
					anchors.fill: parent
					onClicked: { parent.clicked();}
					onEntered: { parent.color = "#eaeaea" }
					onExited: { parent.color = "transparent" }
					onPressed: { parent.color = "#c7c7c7" }
					onReleased: {
						if (containsMouse) {
							parent.color = "#eaeaea"
						} else {
							parent.color = "transparent"
						}
					}
				}
			}

			Rectangle {
				id: refreshButton
				objectName: "refreshButton"
				width: 24
				height: 24
				color: refreshButtonMouseArea.containsMouse ? "#eaeaea" : "transparent"
				signal clicked

				Image {
					anchors.fill: parent
					source: webEngineView.loading ? "qrc:///resources/images/button_stop.png" : "qrc:///resources/images/button_refresh.png"
			 	}

				MouseArea {
					id: refreshButtonMouseArea
					hoverEnabled: true
					anchors.fill: parent
					onClicked: { parent.clicked();}
					onEntered: { parent.color = "#eaeaea" }
					onExited: { parent.color = "transparent" }
					onPressed: { parent.color = "#c7c7c7" }
					onReleased: {
						if (containsMouse) {
							parent.color = "#eaeaea"
						} else {
							parent.color = "transparent"
						}
					}
				}
			}

			Rectangle {
				id: homeButton
				objectName: "homeButton"
				width: 24
				height: 24
				color: homeButtonMouseArea.containsMouse ? "#eaeaea" : "transparent"
				signal clicked

				Image {
					anchors.fill: parent
					source: "qrc:///resources/images/button_home.png"
			 	}

				MouseArea {
					id: homeButtonMouseArea
					hoverEnabled: true
					anchors.fill: parent
					onClicked: { parent.clicked();}
					onEntered: { parent.color = "#eaeaea" }
					onExited: { parent.color = "transparent" }
					onPressed: { parent.color = "#c7c7c7" }
					onReleased: {
						if (containsMouse) {
							parent.color = "#eaeaea"
						} else {
							parent.color = "transparent"
						}
					}
				}
			}

			TextField {
				id: uriField
				objectName: "uriField"
				text: webEngineView.url != "" ? "gopher://" + String(webEngineView.url).replace("http://127.0.0.1:8070/", "") : ""
				placeholderText: ""
				Layout.fillWidth: true
			}

			Button {
				objectName: "goButton"
				text: "Go"
			}
		}

		WebEngineView {
			id: webEngineView
			objectName: "mainView"
            Layout.minimumHeight: 30
            Layout.fillHeight: true
            Layout.fillWidth: true
		}
    }
}
