package utils

import (
	"math/rand"
)

var helloMessages = map[string][]string{
	"en": {"Hi! I can translate your messages to any language.",
		"Hello I can help translate your messages into any language you need.",
		"Hi! I’m here to assist you with translating your messages into any language."},
	"es": {"¡Hola! Puedo traducir tus mensajes a cualquier idioma.",
		"¡Hola! Puedo ayudarte a traducir tus mensajes en cualquier idioma que necesites.",
		"¡Hola! Estoy aquí para ayudarte a traducir tus mensajes a cualquier idioma que necesites."},
	"pt": {"Oi! Eu posso traduzir seus mensagens para qualquer idioma.",
		"Olá! Eu posso ajudar você a traduzir seus mensagens para qualquer idioma.",
		"Olá! Estou aqui para ajudar você a traduzir seus mensagens para qualquer idioma."},
	"fr": {"Salut! Je peux traduire vos messages en tout langage.",
		"Salut! Je peux vous aider à traduire vos messages en tout langage.",
		"Salut! Je suis ici pour vous aider à traduire vos messages en tout langage."},
	"de": {"Hallo! Ich kann Ihre Nachrichten in beliebigen Sprachen übersetzen.",
		"Hallo! Ich kann Ihnen helfen, Ihre Nachrichten in beliebigen Sprachen zu übersetzen.",
		"Hallo! Ich bin hier, um Ihnen zu helfen, Ihre Nachrichten in beliebigen Sprachen zu übersetzen."},
	"it": {"Ciao! Posso tradurre i tuoi messaggi in qualsiasi lingua.",
		"Ciao! Posso aiutarti a tradurre i tuoi messaggi in qualsiasi lingua.",
		"Ciao! Sono qui per aiutarti a tradurre i tuoi messaggi in qualsiasi lingua."},
	"ja": {"こんにちは! 私はあなたのメッセージをどんな言語に翻訳することができます。",
		"こんにちは! 私はあなたにどんな言語に翻訳することができます。",
		"こんにちは！私はあなたのメッセージをどんな言語にも翻訳できます。"},
	"zh": {"你好! 我可以翻译您的消息到任何语言。",
		"你好！我很高兴帮你把消息翻译成任何语言！",
		"你好！我可以帮您将消息翻译成任何语言。"},
	"zh-CN": {"你好! 我可以翻译您的消息到任何语言。",
		"你好！我很高兴帮你把消息翻译成任何语言！",
		"你好！我可以帮您将消息翻译成任何语言。"},
	"id": {"Hai! Saya dapat mengubah pesan Anda ke bahasa apa pun.",
		"Halo! Saya bisa bantu terjemahkan pesan kamu ke bahasa apa aja.",
		"Ciauw! Saya bisa membantu menerjemahkan pesan Anda ke bahasa apa pun."},
	"ru": {"Привет! Я могу переводить ваше сообщения на любой язык.",
		"Привет! Я могу помочь вам переводить сообщения на любой язык.",
		"Привет! Я здесь, чтобы помочь вам переводить сообщения на любой язык."},
	"ar": {"مرحباً! يمكنني ترجمة رسائلك إلى أي لغة.",
		"مرحباً! يمكنني مساعدتك في ترجمة رسائلك إلى أي لغة تحتاجها.",
		"مرحباً! أنا هنا لمساعدتك في ترجمة رسائلك إلى أي لغة."},
	"ko": {"안녕하세요! 저는 당신의 메시지를 어떤 언어로든 번역할 수 있습니다.",
		"안녕하세요! 저는 당신의 메시지를 원하는 언어로 번역하는 것을 도와드릴 수 있습니다.",
		"안녕하세요! 저는 당신의 메시지를 어떤 언어로든 번역하는 데 도움을 드릴 수 있습니다."},
	"pl": {"Cześć! Mogę przetłumaczyć twoje wiadomości na każdy język.",
		"Cześć! Mogę pomóc ci przetłumaczyć twoje wiadomości na dowolny język.",
		"Cześć! Jestem tutaj, aby pomóc ci przetłumaczyć twoje wiadomości na dowolny język."},
	"tr": {"Merhaba! Mesajlarınızı herhangi bir dile çevirebilirim.",
		"Merhaba! Mesajlarınızı ihtiyacınız olan herhangi bir dile çevirmede size yardımcı olabilirim.",
		"Merhaba! Mesajlarınızı herhangi bir dile çevirmede size yardımcı olmak için buradayım."},
	"sv": {"Hej! Jag kan översätta dina meddelanden till vilket språk som helst.",
		"Hej! Jag kan hjälpa dig att översätta dina meddelanden till vilket språk du behöver.",
		"Hej! Jag är här för att hjälpa dig att översätta dina meddelanden till vilket språk som helst."},
	"nl": {"Hoi! Ik kan je berichten naar elke taal vertalen.",
		"Hoi! Ik kan je helpen je berichten naar elke taal te vertalen.",
		"Hoi! Ik ben hier om je te helpen je berichten naar elke taal te vertalen."},
	"da": {"Hej! Jeg kan oversætte dine beskeder til ethvert sprog.",
		"Hej! Jeg kan hjælpe med at oversætte dine beskeder til ethvert sprog.",
		"Hej! Jeg er her for at hjælpe dig med at oversætte dine beskeder til ethvert sprog."},
	"no": {"Hei! Jeg kan oversette meldingene dine til hvilket som helst språk.",
		"Hei! Jeg kan hjelpe deg med å oversette meldingene dine til hvilket som helst språk.",
		"Hei! Jeg er her for å hjelpe deg med å oversette meldingene dine til hvilket som helst språk."},
	"fi": {"Hei! Voin kääntää viestisi mihin tahansa kieleen.",
		"Heippa! Voin auttaa sinua kääntämään viestisi mihin tahansa kieleen.",
		"Heippa! Olen täällä auttamassa viestiesi kääntämisessä mihin tahansa kieleen."},
	"cs": {"Ahoj! Mohu přeložit vaše zprávy do jakéhokoli jazyka.",
		"Ahoj! Mohu vám pomoci přeložit vaše zprávy do jakéhokoli jazyka.",
		"Ahoj! Jsem tady, abych vám pomohl přeložit vaše zprávy do jakéhokoli jazyka."},
	"sk": {"Ahoj! Môžem preložiť vaše správy do akéhokoľvek jazyka.",
		"Ahoj! Môžem vám pomôcť preložiť vaše správy do akéhokoľvek jazyka.",
		"Ahoj! Som tu, aby som vám pomohol preložiť vaše správy do akéhokoľvek jazyka."},
	"hu": {"Helló! Bármilyen nyelvre lefordíthatom az üzeneteidet.",
		"Helló! Segíthetek lefordítani az üzeneteidet bármilyen nyelvre.",
		"Helló! Itt vagyok, hogy segítsek lefordítani az üzeneteidet bármilyen nyelvre."},
	"el": {"Γειά! Μπορώ να μεταφράσω τα μηνύματά σας σε οποιαδήποτε γλώσσα.",
		"Γειά! Μπορώ να σας βοηθήσω να μεταφράσετε τα μηνύματά σας σε οποιαδήποτε γλώσσα.",
		"Γειά! Είμαι εδώ για να σας βοηθήσω να μεταφράσετε τα μηνύματά σας σε οποιαδήποτε γλώσσα."},
	"th": {"สวัสดี! ฉันสามารถแปลข้อความของคุณเป็นภาษาต่างๆ ได้.",
		"สวัสดี! ฉันสามารถช่วยแปลข้อความของคุณเป็นภาษาที่คุณต้องการได้.",
		"สวัสดี! ฉันอยู่ที่นี่เพื่อช่วยแปลข้อความของคุณเป็นภาษาต่างๆ ได้."},
	"hr": {"Bok! Mogu prevesti vaše poruke na bilo koji jezik.",
		"Bok! Mogu vam pomoći prevesti vaše poruke na bilo koji jezik.",
		"Bok! Tu sam da vam pomognem prevesti vaše poruke na bilo koji jezik."},
	"sr": {"Zdravo! Mogu da prevedem vaše poruke na bilo koji jezik.",
		"Zdravo! Mogu vam pomoći da prevedete vaše poruke na bilo koji jezik.",
		"Zdravo! Tu sam da vam pomognem da prevedete vaše poruke na bilo koji jezik."},
	"bg": {"Здравейте! Мога да превеждам вашите съобщения на всякакъв език.",
		"Здравейте! Мога да ви помогна да преведете вашите съобщения на всякакъв език.",
		"Здравейте! Тук съм, за да ви помогна да преведете вашите съобщения на всякакъв език."},
	"ro": {"Bună! Pot traduce mesajele tale în orice limbă.",
		"Bună! Te pot ajuta să îți traduci mesajele în orice limbă.",
		"Bună! Sunt aici să te ajut să îți traduci mesajele în orice limbă."},
	"mk": {"Здраво! Можам да ги преведам вашите пораки на било кој јазик.",
		"Здраво! Можам да ви помогнам да ги преведете вашите пораки на било кој јазик.",
		"Здраво! Тука сум за да ви помогнам да ги преведете вашите пораки на било кој јазик."},
	"sl": {"Pozdrav! Lahko prevedem vaše sporočila v kateri koli jezik.",
		"Pozdrav! Lahko vam pomagam prevesti vaše sporočila v kateri koli jezik.",
		"Pozdrav! Tukaj sem, da vam pomagam prevesti vaše sporočila v kateri koli jezik."},
	"lv": {"Sveiki! Es varu tulkot jūsu ziņas jebkurā valodā.",
		"Sveiki! Es varu jums palīdzēt tulkot jūsu ziņas jebkurā valodā.",
		"Sveiki! Es esmu šeit, lai palīdzētu jums tulkot jūsu ziņas jebkurā valodā."},
	"lt": {"Sveiki! Aš galiu išversti jūsų žinutes į bet kokią kalbą.",
		"Sveiki! Aš galiu jums padėti išversti jūsų žinutes į bet kokią kalbą.",
		"Sveiki! Aš čia, kad jums padėčiau išversti jūsų žinutes į bet kokią kalbą."},
	"et": {"Tere! Ma saan tõlkida teie sõnumeid ükskõik millisesse keelde.",
		"Tere! Ma saan aidata teil tõlkida oma sõnumeid ükskõik millisesse keelde.",
		"Tere! Olen siin, et aidata teil tõlkida oma sõnumeid ükskõik millisesse keelde."},
	"fa": {"سلام! من می‌توانم پیام‌های شما را به هر زبانی ترجمه کنم.",
		"سلام! من می‌توانم به شما در ترجمه پیام‌هایتان به هر زبان کمک کنم.",
		"سلام! من اینجا هستم تا به شما کمک کنم پیام‌هایتان را به هر زبانی ترجمه کنید."},
	"he": {"שלום! אני יכול לתרגם את ההודעות שלך לכל שפה.",
		"שלום! אני יכול לעזור לך לתרגם את ההודעות שלך לכל שפה.",
		"שלום! אני כאן כדי לעזור לך לתרגם את ההודעות שלך לכל שפה."},
	"iw": {"שלום! אני יכול לתרגם את ההודעות שלך לכל שפה.",
		"שלום! אני יכול לעזור לך לתרגם את ההודעות שלך לכל שפה.",
		"שלום! אני כאן כדי לעזור לך לתרגם את ההודעות שלך לכל שפה."},
	"vi": {"Chào! Tôi có thể dịch tin nhắn của bạn sang bất kỳ ngôn ngữ nào.",
		"Chào! Tôi có thể giúp bạn dịch tin nhắn của bạn sang bất kỳ ngôn ngữ nào.",
		"Chào! Tôi ở đây để giúp bạn dịch tin nhắn của bạn sang bất kỳ ngôn ngữ nào."},
	"uk": {"Привіт! Я можу перекласти ваші повідомлення на будь-яку мову.",
		"Привіт! Я можу допомогти вам перекласти ваші повідомлення на будь-яку мову.",
		"Привіт! Я тут, щоб допомогти вам перекласти ваші повідомлення на будь-яку мову."},
	"sq": {"Përshëndetje! Unë mund të përkthej mesazhet tuaja në çdo gjuhë.",
		"Përshëndetje! Unë mund t'ju ndihmoj të përktheni mesazhet tuaja në çdo gjuhë.",
		"Përshëndetje! Jam këtu për t'ju ndihmuar të përktheni mesazhet tuaja në çdo gjuhë."},
	"hy": {"Բարև! Ես կարող եմ ձեր հաղորդագրությունները թարգմանել ցանկացած լեզու.",
		"Բարև! Ես կարող եմ օգնել ձեզ ձեր հաղորդագրությունները թարգմանել ցանկացած լեզու.",
		"Բարև! Ես այստեղ եմ՝ օգնելու ձեզ ձեր հաղորդագրությունները թարգմանել ցանկացած լեզու."},
	"bn": {"হ্যালো! আমি আপনার বার্তা যেকোনো ভাষায় অনুবাদ করতে পারি।",
		"হ্যালো! আমি আপনার বার্তা যেকোনো ভাষায় অনুবাদ করতে সহায়তা করতে পারি।",
		"হ্যালো! আমি এখানে আছি আপনার বার্তা যেকোনো ভাষায় অনুবাদ করতে সহায়তা করার জন্য।"},
	"zu": {"Sawubona! Ngingahumusha imiyalezo yakho noma iyiphi ulimi.",
		"Sawubona! Ngingakusiza ukuhumusha imiyalezo yakho noma iyiphi ulimi.",
		"Sawubona! Ngilapha ukukusiza ukuhumusha imiyalezo yakho noma iyiphi ulimi."},
	"sw": {"Habari! Naweza kutafsiri ujumbe wako kwa lugha yoyote.",
		"Habari! Naweza kukusaidia kutafsiri ujumbe wako kwa lugha yoyote.",
		"Habari! Niko hapa kukusaidia kutafsiri ujumbe wako kwa lugha yoyote."},
	"hi": {"नमस्ते! मैं आपके संदेश को किसी भी भाषा में अनुवाद कर सकता हूँ।",
		"नमस्ते! मैं आपकी मदद कर सकता हूँ आपके संदेश को किसी भी भाषा में अनुवाद करने में।",
		"नमस्ते! मैं यहाँ हूँ आपके संदेश को किसी भी भाषा में अनुवाद करने में आपकी मदद करने के लिए।"},
	"tl": {"Kamusta! Maaari kong isalin ang iyong mga mensahe sa anumang wika.",
		"Kamusta! Maaari kitang tulungan isalin ang iyong mga mensahe sa anumang wika.",
		"Kamusta! Nandito ako upang tulungan kang isalin ang iyong mga mensahe sa anumang wika."},
	"am": {"ሰላም! እባኮትን መልእክቶችዎን በማንኛውም ቋንቋ ማትረት እችላለሁ።",
		"ሰላም! እባኮትን እበልክ መልእክቶችዎን በማንኛውም ቋንቋ ማትረት ማግኘት እችላለሁ።",
		"ሰላም! እባኮትን እኔ እዚህ እስራት መልእክቶችዎን በማንኛውም ቋንቋ ማትረት ለመርዳት እንደ አማራጭ እገኛለሁ።"},
	"af": {"Hallo! Ek kan jou boodskappe na enige taal vertaal.",
		"Hallo! Ek kan jou help om jou boodskappe na enige taal te vertaal.",
		"Hallo! Ek is hier om jou te help om jou boodskappe na enige taal te vertaal."},
}

var instructions = map[string][]string{
	"en": {"Please use the slash command to translate your message.",
		"Kindly use the command to translate your message.",
		"Use the command to translate your message, please."},
	"es": {"Usa el comando para traducir tu mensaje, por fa.",
		"Por favor, usa el comando para traducir tu mensaje.",
		"¡Usa el comando para traducir tu mensaje, vamos!"},
	"fr": {"Veuillez utiliser la commande pour traduire votre message.",
		"Utilisez la commande pour traduire votre message, s'il vous plaît.",
		"Utilisez la commande pour traduire votre message, merci !"},
	"de": {"Bitte verwenden Sie den Befehl, um Ihre Nachricht zu übersetzen.",
		"Nutzen Sie den Befehl, um Ihre Nachricht zu übersetzen, bitte.",
		"Benutzen Sie den Befehl, um Ihre Nachricht zu übersetzen, danke!"},
	"it": {"Per favore usa il comando per tradurre il tuo messaggio.",
		"Usa il comando per tradurre il tuo messaggio, per favore.",
		"Usa il comando per tradurre il tuo messaggio, grazie!"},
	"pt": {"Por favor, use o comando para traduzir sua mensagem.",
		"Use o comando para traduzir sua mensagem, por favor.",
		"Use o comando para traduzir sua mensagem, obrigado!"},
	"ru": {"Пожалуйста, используйте команду для перевода вашего сообщения.",
		"Используйте команду для перевода вашего сообщения, пожалуйста.",
		"Используйте команду для перевода вашего сообщения, спасибо!"},
	"pl": {"Proszę użyj polecenia, aby przetłumaczyć swoją wiadomość.",
		"Użyj polecenia, aby przetłumaczyć swoją wiadomość, proszę.",
		"Użyj polecenia, aby przetłumaczyć swoją wiadomość, dziękuję!"},
	"zh-CN": {"请使用命令来翻译您的消息。",
		"使用命令翻译您的消息，谢谢。",
		"请使用命令翻译您的消息，谢谢！"},
	"ko": {"메시지를 번역하려면 명령어를 사용하세요.",
		"메시지를 번역하려면 명령어를 사용해주세요.",
		"메시지를 번역하려면 명령어를 사용해 주세요!"},
	"ar": {"يرجى استخدام الأمر لترجمة رسالتك.",
		"استخدم الأمر لترجمة رسالتك، من فضلك.",
		"استخدم الأمر لترجمة رسالتك، شكراً!"},
	"ja": {"メッセージを翻訳するにはコマンドを使用してください。",
		"メッセージを翻訳するにはコマンドを使ってください。",
		"メッセージを翻訳するにはコマンドを使ってね！"},
	"hi": {"कृपया अपने संदेश को अनुवादित करने के लिए कमांड का उपयोग करें।",
		"अपने संदेश को अनुवादित करने के लिए कृपया कमांड का उपयोग करें।",
		"कृपया कमांड का उपयोग करके अपने संदेश का अनुवाद करें, धन्यवाद!"},
	"th": {"กรุณาใช้คำสั่งเพื่อแปลข้อความของคุณ.",
		"ใช้คำสั่งเพื่อแปลข้อความของคุณนะครับ/ค่ะ.",
		"กรุณาใช้คำสั่งเพื่อแปลข้อความของคุณ ขอบคุณครับ/ค่ะ!"},
	"tr": {"Lütfen mesajınızı çevirmek için komutu kullanın.",
		"Mesajınızı çevirmek için komutu kullanın, lütfen.",
		"Mesajınızı çevirmek için komutu kullanın, teşekkürler!"},
	"vi": {"Vui lòng sử dụng lệnh để dịch tin nhắn của bạn.",
		"Sử dụng lệnh để dịch tin nhắn của bạn nhé.",
		"Vui lòng dùng lệnh để dịch tin nhắn của bạn, cảm ơn!"},
	"ms": {"Sila gunakan perintah untuk menterjemahkan mesej anda.",
		"Gunakan perintah untuk menterjemahkan mesej anda, sila.",
		"Sila gunakan perintah untuk menterjemahkan mesej anda, terima kasih!"},
	"fa": {"لطفاً از دستور برای ترجمه پیام خود استفاده کنید.",
		"لطفاً از دستور برای ترجمه پیام خود استفاده کنید، با تشکر.",
		"لطفاً از دستور برای ترجمه پیام خود استفاده کنید، مرسی!"},
	"iw": {"בבקשה השתמשו בפקודה כדי לתרגם את ההודעה שלכם.",
		"השתמשו בפקודה כדי לתרגם את ההודעה שלכם, בבקשה.",
		"בבקשה השתמשו בפקודה כדי לתרגם את ההודעה שלכם, תודה!"},
	"sq": {"Ju lutemi përdorni komandën për të përkthyer mesazhin tuaj.",
		"Përdorni komandën për të përkthyer mesazhin tuaj, ju lutem.",
		"Ju lutemi përdorni komandën për të përkthyer mesazhin tuaj, faleminderit!"},
	"cs": {"Použijte příkaz k překladu vaší zprávy.",
		"Použijte příkaz k překladu vaší zprávy, prosím.",
		"Použijte příkaz k překladu vaší zprávy, děkujeme!"},
	"lt": {"Prašome naudoti komandą, kad išverstų jūsų žinutę.",
		"Naudokite komandą, kad išverstų jūsų žinutę, prašome.",
		"Prašome naudoti komandą, kad išverstų jūsų žinutę, ačiū!"},
	"hu": {"Kérlek használd a parancsot az üzeneted lefordításához.",
		"Használd a parancsot az üzeneted lefordításához, kérlek.",
		"Kérlek használd a parancsot az üzeneted lefordításához, köszönöm!"},
	"ro": {"Te rog folosește comanda pentru a traduce mesajul tău.",
		"Folosește comanda pentru a traduce mesajul tău, te rog.",
		"Te rog folosește comanda pentru a traduce mesajul tău, mulțumesc!"},
	"bg": {"Моля, използвайте командата за да преведете съобщението си.",
		"Използвайте командата за да преведете съобщението си, моля.",
		"Моля, използвайте командата за да преведете съобщението си, благодаря!"},
	"sr": {"Molim vas, koristite komandu da prevedete svoju poruku.",
		"Koristite komandu da prevedete svoju poruku, molim vas.",
		"Molim vas, koristite komandu da prevedete svoju poruku, hvala!"},
	"hr": {"Molimo vas, koristite komandu za prevođenje vaše poruke.",
		"Koristite komandu za prevođenje vaše poruke, molimo.",
		"Molimo vas, koristite komandu za prevođenje vaše poruke, hvala!"},
	"sk": {"Prosím použite príkaz na preloženie vašej správy.",
		"Použite príkaz na preloženie vašej správy, prosím.",
		"Prosím použite príkaz na preloženie vašej správy, ďakujem!"},
	"mk": {"Ве молиме користете командата за да го преведете вашето соопштение.",
		"Користете командата за да го преведете вашето соопштение, ве молиме.",
		"Ве молиме користете командата за да го преведете вашето соопштение, благодарам!"},
	"et": {"Palun kasutage käsku oma sõnumi tõlkimiseks.",
		"Kasutage käsku oma sõnumi tõlkimiseks, palun.",
		"Palun kasutage käsku oma sõnumi tõlkimiseks, aitäh!"},
	"sl": {"Prosimo, uporabite ukaz za prevod vaše sporočila.",
		"Uporabite ukaz za prevod vaše sporočila, prosimo.",
		"Prosimo, uporabite ukaz za prevod vaše sporočila, hvala!"},
	"lv": {"Lūdzu, izmantojiet komandu, lai tulkotu jūsu ziņojumu.",
		"Izmantojiet komandu, lai tulkotu jūsu ziņojumu, lūdzu.",
		"Lūdzu, izmantojiet komandu, lai tulkotu jūsu ziņojumu, paldies!"},
	"no": {"Vennligst bruk kommandoen for å oversette meldingen din.",
		"Bruk kommandoen for å oversette meldingen din, vær så snill.",
		"Vennligst bruk kommandoen for å oversette meldingen din, takk!"},
	"da": {"Brug venligst kommandoen til at oversætte din besked.",
		"Brug kommandoen til at oversætte din besked, tak.",
		"Brug venligst kommandoen til at oversætte din besked, mange tak!"},
	"sv": {"Var god använd kommandot för att översätta ditt meddelande.",
		"Använd kommandot för att översätta ditt meddelande, tack.",
		"Var god använd kommandot för att översätta ditt meddelande, tack så mycket!"},
	"fi": {"Käytä komentoa kääntääksesi viestisi.",
		"Please use the command to translate your message.",
		"Käytä komentoa kääntääksesi viestisi, kiitos!"},
	"bs": {"Molimo koristite komandu za prevođenje vaše poruke.",
		"Koristite komandu za prevođenje vaše poruke, molim.",
		"Molimo koristite komandu za prevođenje vaše poruke, hvala!"},
	"gl": {"Por favor, use o comando para traducir a súa mensaxe.",
		"Use o comando para traducir a súa mensaxe, por favor.",
		"Por favor, use o comando para traducir a súa mensaxe, grazas!"},
	"ca": {"Si us plau, utilitzeu la comanda per traduir el vostre missatge.",
		"Utilitzeu la comanda per traduir el vostre missatge, si us plau.",
		"Si us plau, utilitzeu la comanda per traduir el vostre missatge, gràcies!"},
	"he": {"בבקשה השתמש בפקודה כדי לתרגם את ההודעה שלך.",
		"אנא השתמש בפקודה כדי לתרגם את ההודעה שלך.",
		"בבקשה השתמש בפקודה כדי לתרגם את ההודעה שלך, תודה!"},
	"uk": {"Будь ласка, використовуйте команду для перекладу вашого повідомлення.",
		"Використовуйте команду для перекладу вашого повідомлення, будь ласка.",
		"Будь ласка, використовуйте команду для перекладу вашого повідомлення, дякуємо!"},
	"zu": {"Sicela usebenzise umyalo ukuhumusha umlayezo wakho.",
		"Usebenzisa umyalo ukuhumusha umlayezo wakho, sicela.",
		"Sicela usebenzise umyalo ukuhumusha umlayezo wakho, ngiyabonga!"},
	"sw": {"Tafadhali tumia amri kutafsiri ujumbe wako.",
		"Tumia amri kutafsiri ujumbe wako, tafadhali.",
		"Tafadhali tumia amri kutafsiri ujumbe wako, asante!"},
	"tl": {"Pakiusap, gamitin ang command upang isalin ang iyong mensahe.",
		"Gamitin ang command upang isalin ang iyong mensahe, pakiusap.",
		"Pakiusap, gamitin ang command upang isalin ang iyong mensahe, salamat!"},
	"am": {"እባኮትን መልእክትዎን ለመተርጎም ትእዛዝን ተጠቅመው።",
		"እባኮትን መልእክትዎን ለመተርጎም ትእዛዝን ተጠቅመው፣ እናመሰግናለን!"},
	"af": {"Gebruik asseblief die opdrag om jou boodskap te vertaal.",
		"Gebruik die opdrag om jou boodskap te vertaal, asseblief.",
		"Gebruik asseblief die opdrag om jou boodskap te vertaal, dankie!"},
}

func GetRandomGreeting(locale string) string {
	if locale == "" {
		locale = "en"
	}
	greeting := helloMessages[locale]
	if len(greeting) == 0 {
		println(locale + " is not supported, defaulting to 'en' locale")
		greeting = helloMessages["en"]
	}
	return greeting[rand.Intn(len(greeting))]
}

func GetRandomInstruction(locale string) string {
	if locale == "" {
		locale = "en"
	}
	instruction := instructions[locale]
	if len(instruction) == 0 {
		instruction = instructions["en"]
	}
	return instruction[rand.Intn(len(instruction))]
}
