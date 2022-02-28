package main

type Success struct {
	Success bool
}

type Price struct {
	PriceUSD float32 `njson:"collection.payment_tokens.#.usd_price"`
	PriceETH float32 `njson:"collection.payment_tokens.#.eth_price"`
}

type Collection struct {
	Editors       []string `njson:"editors"`
	PaymentTokens []struct {
		ID       int     `njson:"id"`
		Symbol   string  `njson:"symbol"`
		Address  string  `njson:"address"`
		ImageURL string  `njson:"image_url"`
		Name     string  `njson:"name"`
		Decimals int     `njson:"decimals"`
		EthPrice float64 `njson:"eth_price"`
		UsdPrice float64 `njson:"usd_price"`
	} `njson:"payment_tokens"`
	PrimaryAssetContracts []struct {
		Address                     string      `njson:"address"`
		AssetContractType           string      `njson:"asset_contract_type"`
		CreatedDate                 string      `njson:"created_date"`
		Name                        string      `njson:"name"`
		NftVersion                  string      `njson:"nft_version"`
		OpenseaVersion              interface{} `njson:"opensea_version"`
		Owner                       int         `njson:"owner"`
		SchemaName                  string      `njson:"schema_name"`
		Symbol                      string      `njson:"symbol"`
		TotalSupply                 string      `njson:"total_supply"`
		Description                 string      `njson:"description"`
		ExternalLink                string      `njson:"external_link"`
		ImageURL                    string      `njson:"image_url"`
		DefaultToFiat               bool        `njson:"default_to_fiat"`
		DevBuyerFeeBasisPoints      int         `njson:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints     int         `njson:"dev_seller_fee_basis_points"`
		OnlyProxiedTransfers        bool        `njson:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  int         `njson:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints int         `njson:"opensea_seller_fee_basis_points"`
		BuyerFeeBasisPoints         int         `njson:"buyer_fee_basis_points"`
		SellerFeeBasisPoints        int         `njson:"seller_fee_basis_points"`
		PayoutAddress               string      `njson:"payout_address"`
	} `njson:"primary_asset_contracts"`
	Traits struct {
		EHRShares struct {
			Min int `njson:"min"`
			Max int `njson:"max"`
		} `njson:"EHR Shares"`
		RakeBack struct {
			Min int `njson:"min"`
			Max int `njson:"max"`
		} `njson:"Rake Back"`
		Background struct {
			Amethyst    int `njson:"amethyst"`
			Aqua        int `njson:"aqua"`
			Asphalt     int `njson:"asphalt"`
			Bege        int `njson:"bege"`
			Charcoal    int `njson:"charcoal"`
			Coin        int `njson:"coin"`
			Crimson     int `njson:"crimson"`
			DeepPurple  int `njson:"deep purple"`
			Iris        int `njson:"iris"`
			Lemon       int `njson:"lemon"`
			Midnight    int `njson:"midnight"`
			Mint        int `njson:"mint"`
			Navy        int `njson:"navy"`
			Placeholder int `njson:"placeholder"`
			Rose        int `njson:"rose"`
			Rouge       int `njson:"rouge"`
			Saphire     int `njson:"saphire"`
			Sheen       int `njson:"sheen"`
			Sky         int `njson:"sky"`
			Tiger       int `njson:"tiger"`
			Tuscan      int `njson:"tuscan"`
		} `njson:"Background"`
		Clothes struct {
			Eight0SSweater   int `njson:"80s sweater"`
			AceTShirt        int `njson:"ace t-shirt"`
			AllOrNothing     int `njson:"all or nothing"`
			AztecMaxtlatl    int `njson:"aztec maxtlatl"`
			BabyOctopus      int `njson:"baby octopus"`
			BlueRaincoat     int `njson:"blue raincoat"`
			BoxingGloves     int `njson:"boxing gloves"`
			CactusClothes    int `njson:"cactus clothes"`
			CheaterSuite     int `njson:"cheater suite"`
			ChristmasScarf   int `njson:"christmas scarf"`
			Clown            int `njson:"clown"`
			Croupier         int `njson:"croupier"`
			CryptoDancer     int `njson:"crypto dancer"`
			DollarstoreShirt int `njson:"dollarstore shirt"`
			EthCollar        int `njson:"eth collar"`
			EthWings         int `njson:"eth wings"`
			FairyLights      int `njson:"fairy lights"`
			Farmer           int `njson:"farmer"`
			FunkyKid         int `njson:"funky kid"`
			FurOfCards       int `njson:"fur of cards"`
			GamblerSSuite    int `njson:"gambler_s suite"`
			GoldenHakama     int `njson:"golden hakama"`
			GoldenSerpant    int `njson:"golden serpant"`
			GothClothes      int `njson:"goth clothes"`
			GreenJacket      int `njson:"green jacket"`
			GuardSuite       int `njson:"guard suite"`
			HazmatSuite      int `njson:"hazmat suite"`
			HeartOfBtc       int `njson:"heart of btc"`
			Illusionist      int `njson:"illusionist"`
			IrishLuck        int `njson:"irish luck"`
			Leon             int `njson:"leon"`
			LeopardTShirt    int `njson:"leopard t-shirt"`
			LuckyApron       int `njson:"lucky apron"`
			MrSnow           int `njson:"mr. snow"`
			Net              int `njson:"net"`
			NftMan           int `njson:"nft-man"`
			PirateCostume    int `njson:"pirate costume"`
			PizzaRobe        int `njson:"pizza robe"`
			Placeholder      int `njson:"placeholder"`
			PrisonCook       int `njson:"prison cook"`
			PurpleSpot       int `njson:"purple spot"`
			RacingCostume    int `njson:"racing costume"`
			Raver            int `njson:"raver"`
			RedKimono        int `njson:"red kimono"`
			ResponsibleSuit  int `njson:"responsible suit"`
			RoarShirt        int `njson:"roar shirt"`
			RoyalCloak       int `njson:"royal cloak"`
			SafetyVest       int `njson:"safety vest"`
			ScientistRobe    int `njson:"scientist robe"`
			Skeleton         int `njson:"skeleton"`
			StreetBoy        int `njson:"street boy"`
			TexasHoldem      int `njson:"texas holdem"`
			Tomahawk         int `njson:"tomahawk"`
			VikingStrap      int `njson:"viking strap"`
			WildWest         int `njson:"wild west"`
			WinnerSTie       int `njson:"winner_s tie"`
			WitchCloak       int `njson:"witch cloak"`
		} `njson:"Clothes"`
		Ear struct {
			AztecJewellery int `njson:"aztec jewellery"`
			BasicLever     int `njson:"basic lever"`
			BrokenArrow    int `njson:"broken arrow"`
			BronzeBtc      int `njson:"bronze btc"`
			BtcDagger      int `njson:"btc dagger"`
			BtcJackpot     int `njson:"btc jackpot"`
			BtcSafe        int `njson:"btc safe"`
			CactusFlower   int `njson:"cactus flower"`
			ChristmasBalls int `njson:"christmas balls"`
			Cone           int `njson:"cone"`
			DiamondFlowers int `njson:"diamond flowers"`
			DiscoBall      int `njson:"disco ball"`
			DollarRoll     int `njson:"dollar roll"`
			Dynamite       int `njson:"dynamite"`
			EthStick       int `njson:"eth stick"`
			FishFin        int `njson:"fish fin"`
			GoldenArrow    int `njson:"golden arrow"`
			GreenDice      int `njson:"green dice"`
			HiddenMic      int `njson:"hidden mic"`
			KeyOfLife      int `njson:"key of life"`
			NuclearLever   int `njson:"nuclear lever"`
			PaintingBrush  int `njson:"painting brush"`
			Pencil         int `njson:"pencil"`
			PinkAce        int `njson:"pink ace"`
			Placeholder    int `njson:"placeholder"`
			PoisonLiquor   int `njson:"poison liquor"`
			QuadAces       int `njson:"quad aces"`
			Shuriken       int `njson:"shuriken"`
			SnakePin       int `njson:"snake pin"`
			SnowballGlobe  int `njson:"snowball globe"`
			TalismanBone   int `njson:"talisman bone"`
			TopGear        int `njson:"top gear"`
			Treasure       int `njson:"treasure"`
			Watermelon     int `njson:"watermelon"`
			Watts          int `njson:"watts"`
			Wing           int `njson:"wing"`
			WoolBall       int `njson:"wool ball"`
		} `njson:"Ear"`
		Eyes struct {
			Num777         int `njson:"777"`
			SevenBeans     int `njson:"7 beans"`
			AngryStare     int `njson:"angry stare"`
			CandySevens    int `njson:"candy sevens"`
			CashCow        int `njson:"cash cow"`
			CasinoGlasses  int `njson:"casino glasses"`
			ChristmasGifts int `njson:"christmas gifts"`
			DarkMode       int `njson:"dark mode"`
			Dreamer        int `njson:"dreamer"`
			EagleEyes      int `njson:"eagle eyes"`
			EasyLove       int `njson:"easy love"`
			EthGlasses     int `njson:"eth glasses"`
			FlirtyEyes     int `njson:"flirty eyes"`
			FoolSGold      int `njson:"fool_s gold"`
			GearBox        int `njson:"gear box"`
			GoldenLids     int `njson:"golden lids"`
			GothEyes       int `njson:"goth eyes"`
			GreenEyes      int `njson:"green eyes"`
			Killer         int `njson:"killer"`
			LovingGaze     int `njson:"loving gaze"`
			Madman         int `njson:"madman"`
			NuclearEyes    int `njson:"nuclear eyes"`
			OutOfBlue      int `njson:"out of blue"`
			Philosopher    int `njson:"philosopher"`
			Placeholder    int `njson:"placeholder"`
			PurpleBlood    int `njson:"purple blood"`
			RedDiamonds    int `njson:"red diamonds"`
			RedEyes        int `njson:"red eyes"`
			RedMoon        int `njson:"red moon"`
			RoyalStare     int `njson:"royal stare"`
			SamuraiClubs   int `njson:"samurai-clubs"`
			SeaRuler       int `njson:"sea ruler"`
			SerpantEyes    int `njson:"serpant eyes"`
			ShakyBtc       int `njson:"shaky btc"`
			SlotrEyes      int `njson:"slotr eyes"`
			SphinxEyes     int `njson:"sphinx eyes"`
			SpinningClubs  int `njson:"spinning clubs"`
			SquareGlasses  int `njson:"square glasses"`
			StrawberrySpin int `njson:"strawberry spin"`
			Sunglasses     int `njson:"sunglasses"`
			Terminator     int `njson:"terminator"`
			VegasEyes      int `njson:"vegas eyes"`
			Warrior        int `njson:"warrior"`
			Welcome        int `njson:"welcome"`
		} `njson:"Eyes"`
		Hat struct {
			AceHat             int `njson:"ace hat"`
			AceHelmet          int `njson:"ace helmet"`
			BlackHair          int `njson:"black hair"`
			BloodyCards        int `njson:"bloody cards"`
			BrownHair          int `njson:"brown hair"`
			CactusHair         int `njson:"cactus hair"`
			CaptainSHat        int `njson:"captain_s hat"`
			ChristmasTree      int `njson:"christmas tree"`
			CowboyHat          int `njson:"cowboy hat"`
			CreativeIdea       int `njson:"creative idea"`
			DeerHorns          int `njson:"deer horns"`
			DiscoDice          int `njson:"disco dice"`
			Eagle              int `njson:"eagle"`
			Fishbone           int `njson:"fishbone"`
			GoatHorns          int `njson:"goat horns"`
			GorgonSHead        int `njson:"gorgon_s head"`
			IndianPipe         int `njson:"indian pipe"`
			Leaves             int `njson:"leaves"`
			LeopardCap         int `njson:"leopard cap"`
			MegamindBulbs      int `njson:"megamind bulbs"`
			MoctezumaHeaddress int `njson:"moctezuma headdress"`
			Nemes              int `njson:"nemes"`
			Panama             int `njson:"panama"`
			Placeholder        int `njson:"placeholder"`
			PoseidonSCrown     int `njson:"poseidon_s crown"`
			PurpleSkull        int `njson:"purple skull"`
			Riviera            int `njson:"riviera"`
			RoyalCrown         int `njson:"royal crown"`
			SamuraiSHat        int `njson:"samurai_s hat"`
			Shamrock           int `njson:"shamrock"`
			SkullCap           int `njson:"skull cap"`
			SlotieHat          int `njson:"slotie hat"`
			ThanksgivingHat    int `njson:"thanksgiving hat"`
			Uranium            int `njson:"uranium"`
			VikingSHair        int `njson:"viking_s hair"`
			WatermelonCap      int `njson:"watermelon cap"`
			WinnerSHat         int `njson:"winner_s hat"`
		} `njson:"Hat"`
		Mouth struct {
			AztecTongue     int `njson:"aztec tongue"`
			BiteOfGold      int `njson:"bite of gold"`
			BloodyMouth     int `njson:"bloody mouth"`
			BlueChip        int `njson:"blue chip"`
			Bubblegum       int `njson:"bubblegum"`
			Candy           int `njson:"candy"`
			CardTeeth       int `njson:"card teeth"`
			Carrot          int `njson:"carrot"`
			CatSmirk        int `njson:"cat smirk"`
			ChainsawTeeth   int `njson:"chainsaw teeth"`
			ClownNose       int `njson:"clown nose"`
			ClubsMouth      int `njson:"clubs mouth"`
			Displeased      int `njson:"displeased"`
			DynamiteTeeth   int `njson:"dynamite teeth"`
			EasyEight       int `njson:"easy eight"`
			Egg             int `njson:"egg"`
			EthTongue       int `njson:"eth tongue"`
			Facemask        int `njson:"facemask"`
			FelineTusks     int `njson:"feline tusks"`
			FishCookie      int `njson:"fish cookie"`
			FishLips        int `njson:"fish lips"`
			GoldenButton    int `njson:"golden button"`
			GraffitiMask    int `njson:"graffiti mask"`
			GreenCard       int `njson:"green card"`
			GreenTeeth      int `njson:"green teeth"`
			GreenTongue     int `njson:"green tongue"`
			Hazmat          int `njson:"hazmat"`
			HealthyWhite    int `njson:"healthy white"`
			LeopardSmirk    int `njson:"leopard smirk"`
			Lier            int `njson:"lier"`
			Lipstick        int `njson:"lipstick"`
			PharaohPostiche int `njson:"pharaoh postiche"`
			Piercing        int `njson:"piercing"`
			PirateBeard     int `njson:"pirate beard"`
			Placeholder     int `njson:"placeholder"`
			RainbowPot      int `njson:"rainbow pot"`
			Reddish         int `njson:"reddish"`
			Rudolph         int `njson:"rudolph"`
			Samurai         int `njson:"samurai"`
			SeaStar         int `njson:"sea star"`
			SenseiBeard     int `njson:"sensei beard"`
			Serpant         int `njson:"serpant"`
			Smoker          int `njson:"smoker"`
			Spikey          int `njson:"spikey"`
			StarTongue      int `njson:"star tongue"`
			SwingMoustache  int `njson:"swing moustache"`
			Toothpick       int `njson:"toothpick"`
			Turkey          int `njson:"turkey"`
			Tusks           int `njson:"tusks"`
			Vertical        int `njson:"vertical"`
			VikingBeard     int `njson:"viking beard"`
			Vintilator      int `njson:"vintilator"`
			WaterBubbles    int `njson:"water bubbles"`
			ZombieMouth     int `njson:"zombie mouth"`
		} `njson:"Mouth"`
		Skin struct {
			Amethyst       int `njson:"amethyst"`
			AquaGreen      int `njson:"aqua green"`
			AstralStripes  int `njson:"astral stripes"`
			AtlantPink     int `njson:"atlant pink"`
			AztecGray      int `njson:"aztec gray"`
			BlackWood      int `njson:"black wood"`
			BlueSky        int `njson:"blue sky"`
			BrownWood      int `njson:"brown wood"`
			Charcoal       int `njson:"charcoal"`
			Cheeze         int `njson:"cheeze"`
			DiscoGreen     int `njson:"disco green"`
			DiscoPurple    int `njson:"disco purple"`
			FernGreen      int `njson:"fern green"`
			FighterSBronze int `njson:"fighter_s bronze"`
			Golden         int `njson:"golden"`
			GreenCash      int `njson:"green cash"`
			Indian         int `njson:"indian"`
			Leobody        int `njson:"leobody"`
			LuckyBronze    int `njson:"lucky bronze"`
			LuckyTurkey    int `njson:"lucky turkey"`
			MetaViolet     int `njson:"meta violet"`
			NuclearGreen   int `njson:"nuclear green"`
			PersianRed     int `njson:"persian red"`
			PineGreen      int `njson:"pine green"`
			Placeholder    int `njson:"placeholder"`
			PoseidonBlue   int `njson:"poseidon blue"`
			SteelGray      int `njson:"steel gray"`
			Strawberry     int `njson:"strawberry"`
			StreetArt      int `njson:"street art"`
			ThunderGray    int `njson:"thunder gray"`
			Tiger          int `njson:"tiger"`
			Watermelon     int `njson:"watermelon"`
			WornGold       int `njson:"worn gold"`
		} `njson:"Skin"`
	} `njson:"traits"`
	Stats struct {
		OneDayVolume          float64 `njson:"one_day_volume"`
		OneDayChange          float64 `njson:"one_day_change"`
		OneDaySales           float64 `njson:"one_day_sales"`
		OneDayAveragePrice    float64 `njson:"one_day_average_price"`
		SevenDayVolume        float64 `njson:"seven_day_volume"`
		SevenDayChange        float64 `njson:"seven_day_change"`
		SevenDaySales         float64 `njson:"seven_day_sales"`
		SevenDayAveragePrice  float64 `njson:"seven_day_average_price"`
		ThirtyDayVolume       float64 `njson:"thirty_day_volume"`
		ThirtyDayChange       float64 `njson:"thirty_day_change"`
		ThirtyDaySales        float64 `njson:"thirty_day_sales"`
		ThirtyDayAveragePrice float64 `njson:"thirty_day_average_price"`
		TotalVolume           float64 `njson:"total_volume"`
		TotalSales            float64 `njson:"total_sales"`
		TotalSupply           float64 `njson:"total_supply"`
		Count                 float64 `njson:"count"`
		NumOwners             int     `njson:"num_owners"`
		AveragePrice          float64 `njson:"average_price"`
		NumReports            int     `njson:"num_reports"`
		MarketCap             float64 `njson:"market_cap"`
		FloorPrice            float64 `njson:"floor_price"`
	} `njson:"stats"`
	BannerImageURL          string      `njson:"banner_image_url"`
	ChatURL                 interface{} `njson:"chat_url"`
	CreatedDate             string      `njson:"created_date"`
	DefaultToFiat           bool        `njson:"default_to_fiat"`
	Description             string      `njson:"description"`
	DevBuyerFeeBasisPoints  string      `njson:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints string      `njson:"dev_seller_fee_basis_points"`
	DiscordURL              string      `njson:"discord_url"`
	DisplayData             struct {
		CardDisplayStyle string `njson:"card_display_style"`
	} `njson:"display_data"`
	ExternalURL                 string      `njson:"external_url"`
	Featured                    bool        `njson:"featured"`
	FeaturedImageURL            string      `njson:"featured_image_url"`
	Hidden                      bool        `njson:"hidden"`
	SafelistRequestStatus       string      `njson:"safelist_request_status"`
	ImageURL                    string      `njson:"image_url"`
	IsSubjectToWhitelist        bool        `njson:"is_subject_to_whitelist"`
	LargeImageURL               string      `njson:"large_image_url"`
	MediumUsername              interface{} `njson:"medium_username"`
	Name                        string      `njson:"name"`
	OnlyProxiedTransfers        bool        `njson:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string      `njson:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string      `njson:"opensea_seller_fee_basis_points"`
	PayoutAddress               string      `njson:"payout_address"`
	RequireEmail                bool        `njson:"require_email"`
	ShortDescription            interface{} `njson:"short_description"`
	Slug                        string      `njson:"slug"`
	TelegramURL                 interface{} `njson:"telegram_url"`
	TwitterUsername             string      `njson:"twitter_username"`
	InstagramUsername           interface{} `njson:"instagram_username"`
	WikiURL                     interface{} `njson:"wiki_url"`
}
