package instagram

type AutoGenerated struct {
	Graphql struct {
		ShortcodeMedia struct {
			ID         string `json:"id"`
			Shortcode  string `json:"shortcode"` // 連結碼
			Dimensions struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"dimensions"`
			MediaOverlayInfo interface{} `json:"media_overlay_info"`
			MediaPreview     interface{} `json:"media_preview"`
			DisplayURL       string      `json:"display_url"` // 小圖
			DisplayResources []struct {
				Src          string `json:"src"`
				ConfigWidth  int    `json:"config_width"`
				ConfigHeight int    `json:"config_height"`
			} `json:"display_resources"` // 不同解析度
			IsVideo               bool `json:"is_video"`
			EdgeMediaToTaggedUser struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_to_tagged_user"`
			EdgeMediaToCaption struct {
				Edges []struct {
					Node struct {
						Text string `json:"text"` // 內文
						/*
							🔸地瓜麵包 $45 🔸馬鈴薯麵包 $48 - 近期很夯的地瓜麵包你吃過了嗎？不僅外型造型超擬真、超可愛，一口咬下也是濃濃的地瓜香氣，綿密的甜地瓜餡搭配Q軟的麵包外衣，好吃到值得你來排隊～ - 另款馬鈴薯麵包也是還原度十足，內餡是馬鈴薯和莫扎瑞拉起司的組合，吃起來鹹香鹹香的很不錯，六入禮盒還會用地瓜箱子裝，真的吸睛度破表啊！ ⚠️每天14:30出爐，想吃的記得早點來卡位！ - - - - - - - - - - - - - - - - - - - - - 📍 東橋商店 🏠 台北市中山北路一段105巷3之1號 ⏱ 11:00-20:00（週三公休） #popyummy台北 - - - - - - - - - - - - - - - - - - - - - - 📷 @h_a_n.7 Thanks for sharing - 👉 在照片中標注或影片下方 #popyummy 就有機會透過我們將你的照片或影片分享給大家🙆🏻‍♀ - #波波懂吃玩#台北美食#中山美食#中山站美食#東橋商店#麵包#地瓜麵包#馬鈴薯麵包#台北麵包#taipei##taiwan#taipeifood#bread#sweetpotato#potato#yummy#popyummy"
						*/
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_to_caption"`
			HasRankedComments bool `json:"has_ranked_comments"`
			TakenAtTimestamp  int  `json:"taken_at_timestamp"`
			Location          struct {
				ID            string `json:"id"`              // 110196067138447
				HasPublicPage bool   `json:"has_public_page"` //
				Name          string `json:"name"`            // "東橋商店"
				Slug          string `json:"slug"`
				AddressJSON   string `json:"address_json"` // "{"street_address": "\u53f0\u5317\u5e02\u4e2d\u5c71\u5317\u8def\u4e00\u6bb5105\u5df73\u4e4b1\u865f", "zip_code": "104", "city_name": "Taipei, Taiwan", "region_name": "", "country_code": "TW", "exact_city_match": false, "exact_region_match": false, "exact_country_match": false}"
			} `json:"location"`
			Owner struct {
				ID                        string `json:"id"`
				IsVerified                bool   `json:"is_verified"`
				ProfilePicURL             string `json:"profile_pic_url"`
				Username                  string `json:"username"`
				BlockedByViewer           bool   `json:"blocked_by_viewer"`
				RestrictedByViewer        bool   `json:"restricted_by_viewer"`
				FollowedByViewer          bool   `json:"followed_by_viewer"`
				FullName                  string `json:"full_name"`
				HasBlockedViewer          bool   `json:"has_blocked_viewer"`
				IsPrivate                 bool   `json:"is_private"`
				IsUnpublished             bool   `json:"is_unpublished"`
				RequestedByViewer         bool   `json:"requested_by_viewer"`
				PassTieringRecommendation bool   `json:"pass_tiering_recommendation"`
				EdgeOwnerToTimelineMedia  struct {
					Count int `json:"count"`
				} `json:"edge_owner_to_timeline_media"`
				EdgeFollowedBy struct {
					Count int `json:"count"`
				} `json:"edge_followed_by"`
			} `json:"owner"`
			IsAd                       bool `json:"is_ad"`
			EdgeWebMediaToRelatedMedia struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_web_media_to_related_media"`
			EdgeSidecarToChildren struct {
				Edges []struct {
					Node struct { // 其他圖片
						Typename   string `json:"__typename"` // GraphImage, GraphVideo
						ID         string `json:"id"`         // 2419046246401624146
						Shortcode  string `json:"shortcode"`  // CGSLb6UAshS
						Dimensions struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						GatingInfo              interface{} `json:"gating_info"`
						FactCheckOverallRating  interface{} `json:"fact_check_overall_rating"`
						FactCheckInformation    interface{} `json:"fact_check_information"`
						SensitivityFrictionInfo interface{} `json:"sensitivity_friction_info"`
						SharingFrictionInfo     interface{} `json:"sharing_friction_info"`
						MediaOverlayInfo        interface{} `json:"media_overlay_info"`
						MediaPreview            string      `json:"media_preview"`
						DisplayURL              string      `json:"display_url"` // 圖片網址
						DisplayResources        []struct {
							Src          string `json:"src"`
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
						} `json:"display_resources"`
						AccessibilityCaption  string `json:"accessibility_caption"`
						IsVideo               bool   `json:"is_video"`
						TrackingToken         string `json:"tracking_token"`
						EdgeMediaToTaggedUser struct {
							Edges []interface{} `json:"edges"`
						} `json:"edge_media_to_tagged_user"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_sidecar_to_children"`
		} `json:"shortcode_media"`
	} `json:"graphql"`
}