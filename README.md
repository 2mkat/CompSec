# CompSec
Статья.
1) Какую проблему решает эта статья?

Использование паролей и необходимость их защиты никуда не делись. Большинство веб-сайтов, требующих аутентификации, продолжают поддерживать аутентификацию по паролю. Даже высокозащищенные приложения, такие как порталы интернет-банкинга, которые развертывают 2-факторную аутентификацию, полагаются на аутентификацию паролем в качестве одного из факторов аутентификации. Однако фишинговые атаки продолжают мешать аутентификации на основе паролей, несмотря на агрессивные усилия по обнаружению и уничтожению, а также комплексные программы повышения осведомленности пользователей и обучения. В настоящее время нет надежного механизма даже для веб-сайтов, заботящихся о безопасности, чтобы предотвратить перенаправление пользователей на мошеннические веб-сайты и фишинг их паролей.
В этой статье мы применяем анализ угроз в процессе создания веб-пароля и обнаруживаем уязвимость дизайна в поле HTML <input type="password">. Эта уязвимость может быть использована для фишинговых атак, поскольку процесс веб-аутентификации не полностью защищен от каждого поля ввода пароля на веб-сервере. В статье определяются четыре свойства, которые инкапсулируют требования для предотвращения веб-фишинга паролей, и предлагается безопасный протокол для использования с новым полем учетных данных, соответствующим этим четырем свойствам.

•	Идентификация — это заявление о том, кем вы являетесь. В зависимости от ситуации, это может быть имя, адрес электронной почты, номер учетной записи, итд.
•	Аутентификация — предоставление доказательств, что вы на самом деле есть тот, кем идентифицировались (от слова “authentic” — истинный, подлинный).
•	Авторизация — проверка, что вам разрешен доступ к запрашиваемому ресурсу.
•	Фишинг - Злоумышленник копирует исходный код страницы, к примеру, авторизации почтового сервиса, и загружает на свой арендованный хостинг, где он, само собой, разместил свои фиктивные данные. Затем он создает адрес этой страницы очень похожий на оригинальный.Также фиктивная страница настроена так, что после ввода данных (логина и пароля) они сохраняются на сайте хакера. Этот метод основывается на том, что пользователь должен предоставить username и password для успешной идентификации и аутентификации в системе. Пара username/password задается пользователем при его регистрации в системе, при этом в качестве username может выступать адрес электронной почты пользователя.

2) Как определяется безопасность? Какой же противник предполагается?

Противник, с которым мы имеем дело, - это фишинговый злоумышленник с целью успешной аутентификации на честном веб-сервере от имени пользователя (то есть, в частности, противник заинтересован в краже учетных данных пользователей). Мы предполагаем, что противник способен:
- Обман пользователя, чтобы посетить фишинговый веб-сервер под контролем противника.
- Считывает все данные передачи, отправленные из браузера пользователя.
- Прочитайте все данные, хранящиеся в хранилище учетных данных честного веб-сервера.
Далее мы предполагаем, что противник не заинтересован в атаках типа "отказ в обслуживании"; программное обеспечение браузера, работающее на устройстве пользователя, не скомпрометировано; сеанс входа в систему проходит по защищенному каналу (например, TLS), и противник не имеет доступа на запись к честному веб-серверу или хранилищу учетных данных.
