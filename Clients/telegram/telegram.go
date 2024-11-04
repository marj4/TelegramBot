package telegram

import (
	"bot/lid/e"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string //Sample host: tg-bot.com
	basePath string //Sample basepath: /bot<token>
	// tg-bot.com/bot<token>
	client http.Client
}

// Function for create client
func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

//Клиент будет выполнять 2 метода Update(получение новых сообщений) и GiveUpdate(отправка своих собственных)

func (c *Client) Update(offset int, limit int) ([]Update, error) {
	//Метод будет возвращать информацию о сообщении
	//Создадем файл, где будут все структуры типов ./telegram/types.go

	//Формируем параметры запроса
	q := url.Values{}

	//Добавляем параметры к запросу
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	//Отправляем запрос
	//Так код для отправки запроса будет одинаковым для всех методов клиента, то я вынесу его в виде функции

}

// Функция для отправки запроса
func (c *Client) doRequest(methood string, query url.Values) (data []byte, err error) {
	//What is structure defer(read about this)

	const errMsg = "cant do request:"

	//Формируем url на который надо отправить запрос
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, methood),
	} //Как я понял url будет таким: /bot<token>/method(getUpdate)

	//Формируем объект запроса(подготавлием запрос к отправке)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	//Теперь надо передать объект запроса(req) в параметр запроса(query url.Values)
	req.URL.RawQuery = query.Encode() //Метод Encode будет приведёт запрос в тип, который принимает сервер

	//Отправляем, получившееся запрос
	resp, err := c.client.Do(req)
	if err != nil {
		//Так как подобное 'оборачивание' ошибки будет встречаться не раз,создадим пакет Error
		//Путь ./lib/e/e.go
		return nil, e.Wrap(errMsg, err)
	}

}
