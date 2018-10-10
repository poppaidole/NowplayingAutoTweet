package main

import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
 
  "os/exec"
  "time"
)
func errCheck(er error){
  if er != nil {
    panic(er)
  }
}
func check(st string)string{
  newSt := fmt.Sprintf("#NowPlaying: %s",getTrackInfo())
  if st != newSt{
    post(newSt)
    return newSt
  }
  return st
}
//supervisor
func GetTwitterAi() *anaconda.TwitterApi {
  anaconda.SetConsumerKey("")
  anaconda.SetConsumerSecret("")
  api := anaconda.NewTwitterApi("", "")
  return api
}
func getTrackInfo()string{
  infoAscii, err := exec.Command("osascript", "getTrack.scpt").Output()
  errCheck(err)
  info:=string(infoAscii)
  return info
}
func post(text string){
  api :=GetTwitterApi()
  _, err := api.PostTweet(text,nil)
  errCheck(err)
}
func main(){
  tweet := fmt.Sprintf("#NowPlaying: %s",getTrackInfo())
  for{
    tweet =check(tweet)
    time.Sleep(1)
  }
}
