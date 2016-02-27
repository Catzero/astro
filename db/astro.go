package db

type AstroFortune struct {
	Astro     string `json:Astro`
	Date      string `json:Date`
	TodayDesc string `json:TodayDesc`
	WeekDesc  string `json:WeekDesc`
}

type AstroTodayDesc struct {
	Astro   string `json:Astro`
	Chinese string `json:Chinese`
	Zonghe  string `json:Zonghe`
	Love    string `json:Love`
	Work    string `json:Work`
	Money   string `json:Money`
	Health  string `json:Health`
	Kaiyun  string `json:Kaiyun`
	Color   string `json:Color`
	Number  string `json:Number`
	Friend  string `json:Friend`
	Detail  string `json:Detail`
	Time    string `json:Time`
}

type AstroWeekDesc struct {
	Astro    string `json:Astro`
	General  string `json:General`
	Love     string `json:Love`
	Work     string `json:Work`
	RedAlarm string `json:RedAlarm`
	Lucky    string `json:Lucky`
	Position string `json:Position`
	Try      string `json:Try`
	Friend   string `json:Friend`
	Range    string `json:Range`
}
