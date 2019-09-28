package main

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/mitchellh/go-homedir"
	"google.golang.org/api/option"
	"log"
	"os/exec"
	"time"
)

// firebase daemon?
func aaa() {
	ctx := context.Background()
	d,_:=homedir.Dir()
	opt := option.WithCredentialsFile(d+"\\pridaemonchan-firebase-adminsdk-xwp84-e05a47677a.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
		return
	}
	db, e1 := app.DatabaseWithURL(ctx,"https://pridaemonchan.firebaseio.com/")
	if e1 !=nil{
		panic(e1)
	}

	var v string
	ref:=db.NewRef("/test")
	etag,e2 := ref.GetWithETag(ctx, &v)
	if e2!=nil{
		panic(e2)
	}
	for  {
		var change bool
		change,etag,e2= ref.GetIfChanged(ctx,etag, &v)
		if e2!=nil{
			panic(e2)
		}
		if change{
			lunch(string(v))
		}
		time.Sleep(1*time.Second)

	}
}
func lunch(str string){
	log.Printf("%#v",str)
	err := exec.Command("C:\\Program Files\\VideoLAN\\VLC\\vlc.exe",
		"E:\\ダウンロード\\goomalling-storm.mp4",
		"vlc://quit",
		"-f",
		).Run()
	if(err!=nil){
		log.Fatalf("error",err)
	}
}

func main() {
	aaa()

}
