package translations

func initLithuanianTranslation() Translation {
	translation := createTranslation()

	translation.put("requires-js", "Šiam puslapiui reikalingas Javascript jog veiktų teisingai")

	translation.put("start-the-game", "Pradėti žaidimą")
	translation.put("start", "Pradėti")
	translation.put("game-not-started-title", "Žaidimas dar neprasidėjo")
	translation.put("waiting-for-host-to-start", "Palaukite kol organizatorius paleis žaidimą")

	translation.put("round", "Lygis")
	translation.put("toggle-soundeffects", "Garso efektai")
	translation.put("change-your-name", "Pakeisti vardą")
	translation.put("randomize", "Atsitiktinis")
	translation.put("apply", "Patvirtinti")
	translation.put("save", "Išsaugoti")
	translation.put("votekick-a-player", "Balsuoti išspirti žaidėją")
	translation.put("time-left", "Laikas")

	translation.put("last-turn", "(Paskutinis ėjimas: %s)")

	translation.put("drawer-kicked", "Kadangi piešiantis žaidėjas buvo išspirtas, niekas negauna taškų už šį raundą")
	translation.put("self-kicked", "Jūs buvote išspirtas")
	translation.put("kick-vote", "(%s/%s) žaidėjai balsvao išspirti %s.")
	translation.put("player-kicked", "Žaidėjas buvo išspirtas.")
	translation.put("owner-change", "%s yra naujasis vedėjas")

	translation.put("change-lobby-settings", "Keisti nustatymus")
	translation.put("lobby-settings-changed", "Nustatymai pakeisti")
	translation.put("advanced-settings", "Papildomi nustatymai")
	translation.put("word-language", "Žodyno kalba")
	translation.put("drawing-time-setting", "Piešimo laikas")
	translation.put("rounds-setting", "Raundai")
	translation.put("max-players-setting", "Maximalus žaidėjų skaičius")
	translation.put("public-lobby-setting", "Viešas kambarys")
	translation.put("custom-words", "Papildomi žodžiai")
	translation.put("custom-words-info", "Įveskite papildomus žodžius atskirdami juos kabutėmis")
	translation.put("custom-words-chance-setting", "Papildomų žodžių šansas")
	translation.put("players-per-ip-limit-setting", "Žaidėjų iš vieno IP adreso limitas")
	translation.put("enable-votekick-setting", "Leisti balsavimą išspirti")
	translation.put("save-settings", "Išsaugoti nustatymus")
	translation.put("input-contains-invalid-data", "Jūsų pateikti duomenys netinkami")
	translation.put("please-fix-invalid-input", "Pataisykite duomenis ir bandykte iš naujo")
	translation.put("create-lobby", "Sukurti kambarį")

	translation.put("players", "Žaidėjai")
	translation.put("refresh", "Atnaujinti")
	translation.put("join-lobby", "Prisijungti prie kambario")

	translation.put("message-input-placeholder", "Spėjimus ir žinutes rašykite čia")

	translation.put("choose-a-word", "Pasirinkite žodį")
	translation.put("waiting-for-word-selection", "Laukiama žodžio pasirinkimo")
	//This one doesn't use %s, since we want to make one part bold.
	translation.put("is-choosing-word", "renkasi žodį.")

	translation.put("close-guess", "'%s' beveik pataikė.")
	translation.put("correct-guess", "Jūs teisingai atspėjote žodį.")
	translation.put("correct-guess-other-player", "%s teisingai atspėjo žodį.")
	translation.put("round-over", "Ėjimas baigtas, nebuvo pasirinktas žodis.")
	translation.put("round-over-no-word", "Ėjimas baigtas, žodis buvo '%s'.")
	translation.put("game-over-win", "Sveikiname. Jūs laimėjote!")
	translation.put("game-over", "Jūsu vieta %s. su %s taškais")

	translation.put("change-active-color", "Pakeisti aktyvią spalvą")
	translation.put("use-pencil", "Naudoti pieštuką")
	translation.put("use-eraser", "Naudoti trintuką")
	translation.put("use-fill-bucket", "Naudoti kibirą (Užpildo pasirinktą erdvę pasirinkta spalva)")
	translation.put("change-pencil-size-to", "Pakeisti pieštuką / trintuko dydis į %s")
	translation.put("clear-canvas", "Išvalyti paveikslą")

	translation.put("connection-lost", "Nutrūko ryšys!")
	translation.put("connection-lost-text", "Bandoma sugrįžti"+
		" ...\n\nPatikrinkite ar yra interneto ryšys.\nJeigu "+
		"problema tesiasi, susisiekit su puslapio kurėju.")
	translation.put("error-connecting", "Klaida jungiantis prie serverio")
	translation.put("error-connecting-text",
		"Scribble.rs couldn't establish a socket connection.\n\nNors jūsų interneto ryšys "+
			"atrodo veikiantis, arba \npuslapis ar ugniasienė "+
			"yra netesingai nustatyta.\n\nPerkraukite puslapį ir bandykite dar kartą.")

	//Generic words
	//As "close" in "closing the window"
	translation.put("close", "Uždaryti")
	translation.put("no", "Ne")
	translation.put("yes", "Taip")
	translation.put("system", "Sistema")

	translation.put("source-code", "Source Code")
	translation.put("help", "Pagalba")
	translation.put("contact", "Susisiekti")
	translation.put("submit-feedback", "Parašyti atsiliepimą")
	translation.put("stats", "Statusas")

	RegisterTranslation("lt", translation)
	RegisterTranslation("lt-LT", translation)

	return translation
}
