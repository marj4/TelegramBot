package telegram

import (
	"bot/lid/e"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdateMethood   = "getUpdate"
	sendMessageMethood = "sendMessage"
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

//Клиент будет выполнять 2 метода Update(получение новых сообщений) и sendMessage(отправка своих собственных сообщений)

// Метод выполняет запрос к телеграм API, чтобы получить новые сообщения(обновления)
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
	data, err := c.doRequest(getUpdateMethood, q)
	if err != nil {
		return nil, err
	}

	//Я получил данные из ответа(например, какое-то сообщение)
	//Знаю, что там будут json файлы. Объявляю переменную типа UpdatesResponse
	var res UpdatesResponse
	//Парсим данные из data в res(указывая его адрес в памяти)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	//Возвращаем результат
	return res.Result, err

}

// Метод выполнения отправки сообщения
func (c *Client) SendMessage(chatId int, text string) error {
	q := url.Values{}

	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethood, q)
	if err != nil {
		return e.WrapIfErr("cant send message", err)
	}

	return nil
}

// Функция для отправки запроса
func (c *Client) doRequest(methood string, query url.Values) (data []byte, err error) {
	//What is structure defer(read about this)
	defer func() { err = e.WrapIfErr("cant do request", err) }()

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
		return nil, err
	}

	//Теперь надо передать объект запроса(req) в параметр запроса(query url.Values)
	req.URL.RawQuery = query.Encode() //Метод Encode будет приведить запрос в тип, который принимает сервер

	//Отправляем, получившееся запрос
	resp, err := c.client.Do(req)
	if err != nil {
		//Так как подобное 'оборачивание' ошибки будет встречаться не раз,создадим пакет Error
		//Путь ./lib/e/e.go
		return nil, err
	}
	//Close the body response
	defer func() { _ = resp.Body.Close() }()

	//Получаем содержимое тела
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//Возвращаем результат
	return body, nil
}
