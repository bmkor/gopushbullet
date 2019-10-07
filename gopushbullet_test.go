package gopushbullet

import "testing"

func TestHandypushbullet(t *testing.T) {
	pbn, _ := New("config/pushbullet.yml")
	t.Log(pbn.config)
	pbn.Notify("title", "testing")
}
