package timer

import "time"

const poll_rate = 10*time.Millisecond

type timerstruct struct{
  start int64
  runtime int64
}

func init_timer(time_ms int64, t* timerstruct){
  t.start = time.Now().UnixNano() / int64(time.Millisecond)
	t.runtime = time_ms
}
func reset(t* timerstruct){
  t.start = time.Now().UnixNano() / int64(time.Millisecond)
}
func timeout(t* timerstruct) bool {
  if((t.start + t.runtime) < (time.Now().UnixNano() / int64(time.Millisecond))){
    return true
  }
  return false
}
func poll_timer(receiver chan<- bool, t* timerstruct){
	for {
		time.Sleep(poll_rate)
		verify := timeout(t)
		if verify == true {
			receiver <- true
		}
	}
}
