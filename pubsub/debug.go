package pubsub

import "fmt"
import "sort"

type sortResults []ConnectResult

func (x sortResults) Len() int { return len(x) }
func (x sortResults) Less(i, j int) bool {
	itime := x[i].StartTime
	jtime := x[j].StartTime
	dtime := jtime.Sub(itime)
	return dtime > 0
	//return itime.Nanosecond() < jtime.Nanosecond()
}
func (x sortResults) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// CDebug is connect debug.
func CDebug(cResults []ConnectResult) {

	sort.Sort(sortResults(cResults))

	for _, r := range cResults {
		//test := r.StartTime.Nanosecond
		fmt.Printf("ID=%s, sTime=%s, eTime=%s, Durtime=%s \n",
			r.ClientID, r.StartTime, r.EndTime, r.DurTime)
	}

	fastTime := cResults[0].StartTime
	slowTime := cResults[len(cResults)-1].StartTime
	durtime := slowTime.Sub(fastTime)
	clientNum := int64(len(cResults))

	nanoTime := durtime.Nanoseconds()                 //nano秒に変換
	perClient := nanoTime / clientNum                 //1connectにかかったnano秒
	throuput := float64(1000000 / float64(perClient)) //1ms=1000000. 1ms/コネクション時間

	fmt.Printf("#### dtime= %s, clientNum=%d, duration=%dns, %dns/clinet, %f client/ms #### \n",
		durtime, clientNum, nanoTime, perClient, throuput)
}
