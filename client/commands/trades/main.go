package TradesCommand

type Trade struct {
	SecId      int     `xml:"secid"`
	TradeNo    int64   `xml:"tradeno"`
	OrderNo    int64   `xml:"orderno"`
	Board      string  `xml:"board"`
	SecCode    string  `xml:"seccode"`
	Client     string  `xml:"client"`
	Union      string  `xml:"union"`
	BuySell    string  `xml:"buysell"`          // B - покупка, S - продажа
	Time       string  `xml:"time"`             // время сделки
	BrokerRef  string  `xml:"brokerref"`        // примечание
	Value      float64 `xml:"value"`            // объем сделки
	Comission  float64 `xml:"comission"`        // комиссия
	Price      float64 `xml:"price"`            // цена
	Items      int64   `xml:"items"`            // кол-во инструмента в сделках в штуках
	Quantity   int64   `xml:"quantity"`         // количество лотов
	Yield      float64 `xml:"yield"`            // доходность
	Accruedint float64 `xml:"accruedint"`       // НКД
	TradeType  string  `xml:"tradetype"`        // тип сделки: ‘T’ – обычная, ‘N’ – РПС, ‘R’ – РЕПО, ‘P’ – размещение
	SettleCode string  `xml:"settlecode"`       // код поставки
	CurrentPos string  `xml:"currentpos"`       // Текущая позиция по инструменту
	Bypass     int     `xml:"bypass,omitempty"` // Признак внебиржевой сделки, 1 или поле отсутсвует
	Venue      string  `xml:"venue"`            // Площадка (execution place)
}

// Сделки
type Trades struct {
	Items []Trade `xml:"trade"`
}
