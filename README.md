# GoRSSE

Go Remote Server Side Events - Multiple SSE Server, which allows sending messages through RPC.
Позволяет отправлять сообщения веб-приложенияем, основанным на микросервисной архитектуре.

## Usage
1. Создаём `server.go`. В нём определяем серверы SSE в массиве объектов SSE.SSEServerSettings. Для каждого сервера нужно указать имя и порт.
Также нужно указать адрес сервера RPC и его порт. Адрес сервера RPC будет использоваться как адрес для каждого сервера SSE.
*server.go*
    ```
    package main

    func main() {
        //Массив настроек для каждого сервера SSE
        settings := []SSE.SSEServerSettings{{AppName: "SSE1", Port: 3001}, {AppName: "SSE2", Port: 3002}}
        //Старт всех серверов SSE и сервера RPC
        GoRSSE.StartServer("localhost", 3000, settings)
    }
    ```

2. Вызываем метод GoRSSE.SendMsg("адрес_RPC:порт_RPC", "Имя_сервера_SSE", "сообщение").
Данный метод нужно вызывать на серверной стороне веб-приложения, чтобы отправить сообщение нужным клиентам.
    ```
    GoRSSE.SendMsg("localhost:3000", "SSE1", "test message for SSE1")
    ```

3. На стороне клиента в `JavaScript` создаём объект `EventSource` с указанием необходимого сервера SSE и определяем методы для перехвата событий EventSource
    ```
    let eventSourse = new EventSource("http://localhost:3001");
    eventSourse.onmessage = (e)=>{
       console.log(e.data);
    }
    ```

4. После запуска сервера из пунта 1 при вызове метода из пункта 2 на всех клиентах, подключенных к localhost:3001 в консоли появится сообщение
    ```
    test message for SSE1
    ```
