package pubsub

import "fmt"
import "sort"

type sortResults []ConnectResult

func (x sortResults) Len() int { return len(x) }
func (x sortResults) Less(i, j int) bool {
	itime := x[i].StartTime
	jtime := x[j].StartTime
	return itime.Nanosecond() < jtime.Nanosecond()
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
	dtime := slowTime.Sub(fastTime)
	clientNum := len(cResults)

	fmt.Printf("#### dtime= %s, clientNum=%d #### \n", dtime, clientNum)
}
