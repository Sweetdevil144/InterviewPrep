package main

type Val struct {
	value  string
	timestamp int
}

type TimeMap struct {
	data map[string][]Val
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]Val),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], Val{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	vals, ok := this.data[key]
	if !ok {
		return ""
	}
	var res string
	l, r := 0, len(vals)-1
	for l <= r {
		m := l + (r-l)/2
		if vals[m].timestamp <= timestamp {
			res = vals[m].value
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
