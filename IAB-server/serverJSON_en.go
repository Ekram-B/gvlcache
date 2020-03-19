package main

func generatePrettifiedOutPutEN() string {
	return `{
		"gvlSpecificationVersion": 2,
		"vendorListVersion": 29,
		"tcfPolicyVersion": 2,
		"lastUpdated": "2020-03-12T16:05:14Z",
		"purposes": {
		  "1": {
			"id": 1,
			"name": "Store and/or access information on a device",
			"description": "Cookies, device identifiers, or other information can be stored or accessed on your device for the purposes presented to you.",
			"descriptionLegal": "Vendors can:\n* Store and access information on the device such as cookies and device identifiers presented to a user."
		  },
		  "2": {
			"id": 2,
			"name": "Select basic ads",
			"description": "Ads can be shown to you based on the content you’re viewing, the app you’re using, your approximate location, or your device type.",
			"descriptionLegal": "To do basic ad selection vendors can:\n* Use real-time information about the context in which the ad will be shown, to show the ad, including information about the content and the device, such as: device type and capabilities, user agent, URL, IP address\n* Use a user’s non-precise geolocation data\n* Control the frequency of ads shown to a user.\n* Sequence the order in which ads are shown to a user.\n* Prevent an ad from serving in an unsuitable editorial (brand-unsafe) context\nVendors cannot:\n* Create a personalised ads profile using this information for the selection of future ads.\n* N.B. Non-precise means only an approximate location involving at least a radius of 500 meters is permitted."
		  },
		  "3": {
			"id": 3,
			"name": "Create a personalised ads profile",
			"description": "A profile can be built about you and your interests to show you personalised ads that are relevant to you.",
			"descriptionLegal": "To create a personalised ads profile vendors can:\n* Collect information about a user, including a user's activity, interests, demographic information, or location, to create or edit a user profile for use in personalised advertising.\n* Combine this information with other information previously collected, including from across websites and apps, to create or edit a user profile for use in personalised advertising."
		  },
		  "4": {
			"id": 4,
			"name": "Select personalised ads",
			"description": "Personalised ads can be shown to you based on a profile about you.",
			"descriptionLegal": "To select personalised ads vendors can:\n* Select personalised ads based on a user profile or other historical user data, including a user’s prior activity, interests, visits to sites or apps, location, or demographic information."
		  },
		  "5": {
			"id": 5,
			"name": "Create a personalised content profile",
			"description": "A profile can be built about you and your interests to show you personalised content that is relevant to you.",
			"descriptionLegal": "To create a personalised content profile vendors can:\n* Collect information about a user, including a user's activity, interests, visits to sites or apps, demographic information, or location, to create or edit a user profile for personalising content.\n* Combine this information with other information previously collected, including from across websites and apps, to create or edit a user profile for use in personalising content."
		  },
		  "6": {
			"id": 6,
			"name": "Select personalised content",
			"description": "Personalised content can be shown to you based on a profile about you.",
			"descriptionLegal": "To select personalised content vendors can:\n* Select personalised content based on a user profile or other historical user data, including a user’s prior activity, interests, visits to sites or apps, location, or demographic information."
		  },
		  "7": {
			"id": 7,
			"name": "Measure ad performance",
			"description": "The performance and effectiveness of ads that you see or interact with can be measured.",
			"descriptionLegal": "To measure ad performance vendors can:\n* Measure whether and how ads were delivered to and interacted with by a user\n* Provide reporting about ads including their effectiveness and performance\n* Provide reporting about users who interacted with ads using data observed during the course of the user's interaction with that ad\n* Provide reporting to publishers about the ads displayed on their property\n* Measure whether an ad is serving in a suitable editorial environment (brand-safe) context\n* Determine the percentage of the ad that had the opportunity to be seen and the duration of that opportunity\n* Combine this information with other information previously collected, including from across websites and apps\nVendors cannot:\n*Apply panel- or similarly-derived audience insights data to ad measurement data without a Legal Basis to apply market research to generate audience insights (Purpose 9)"
		  },
		  "8": {
			"id": 8,
			"name": "Measure content performance",
			"description": "The performance and effectiveness of content that you see or interact with can be measured.",
			"descriptionLegal": "To measure content performance vendors can:\n* Measure and report on how content was delivered to and interacted with by users.\n* Provide reporting, using directly measurable or known information, about users who interacted with the content\n* Combine this information with other information previously collected, including from across websites and apps.\nVendors cannot:\n* Measure whether and how ads (including native ads) were delivered to and interacted with by a user.\n* Apply panel- or similarly derived audience insights data to ad measurement data without a Legal Basis to apply market research to generate audience insights (Purpose 9)"
		  },
		  "9": {
			"id": 9,
			"name": "Apply market research to generate audience insights",
			"description": "Market research can be used to learn more about the audiences who visit sites/apps and view ads.",
			"descriptionLegal": "To apply market research to generate audience insights vendors can:\n* Provide aggregate reporting to advertisers or their representatives about the audiences reached by their ads, through panel-based and similarly derived insights.\n* Provide aggregate reporting to publishers about the audiences that were served or interacted with content and/or ads on their property by applying panel-based and similarly derived insights.\n* Associate offline data with an online user for the purposes of market research to generate audience insights if vendors have declared to match and combine offline data sources (Feature 1)\n* Combine this information with other information previously collected including from across websites and apps. \nVendors cannot:\n* Measure the performance and effectiveness of ads that a specific user was served or interacted with, without a Legal Basis to measure ad performance.\n* Measure which content a specific user was served and how they interacted with it, without a Legal Basis to measure content performance."
		  },
		  "10": {
			"id": 10,
			"name": "Develop and improve products",
			"description": "Your data can be used to improve existing systems and software, and to develop new products",
			"descriptionLegal": "To develop new products and improve products vendors can:\n* Use information to improve their existing products with new features and to develop new products\n* Create new models and algorithms through machine learning\nVendors cannot:\n* Conduct any other data processing operation allowed under a different purpose under this purpose"
		  }
		},
		"specialPurposes": {
		  "1": {
			"id": 1,
			"name": "Ensure security, prevent fraud, and debug",
			"description": "Your data can be used to monitor for and prevent fraudulent activity, and ensure systems and processes work properly and securely.",
			"descriptionLegal": "To ensure security, prevent fraud and debug vendors can:\n* Ensure data are securely transmitted\n* Detect and prevent malicious, fraudulent, invalid, or illegal activity.\n* Ensure correct and efficient operation of systems and processes, including to monitor and enhance the performance of systems and processes engaged in permitted purposes\nVendors cannot:\n* Conduct any other data processing operation allowed under a different purpose under this purpose."
		  },
		  "2": {
			"id": 2,
			"name": "Technically deliver ads or content",
			"description": "Your device can receive and send information that allows you to see and interact with ads and content.",
			"descriptionLegal": "To deliver information and respond to technical requests vendors can:\n* Use a user’s IP address to deliver an ad over the internet\n* Respond to a user’s interaction with an ad by sending the user to a landing page\n* Use a user’s IP address to deliver content over the internet\n* Respond to a user’s interaction with content by sending the user to a landing page\n* Use information about the device type and capabilities for delivering ads or content, for example, to deliver the right size ad creative or video file in a format supported by the device\nVendors cannot:\n* Conduct any other data processing operation allowed under a different purpose under this purpose"
		  }
		},
		"features": {
		  "1": {
			"id": 1,
			"name": "Match and combine offline data sources",
			"description": "Data from offline data sources can be combined with your online activity in support of one or more purposes",
			"descriptionLegal": "Vendors can:\n* Combine data obtained offline with data collected online in support of one or more Purposes or Special Purposes."
		  },
		  "2": {
			"id": 2,
			"name": "Link different devices",
			"description": "Different devices can be determined as belonging to you or your household in support of one or more of purposes.",
			"descriptionLegal": "Vendors can:\n* Deterministically determine that two or more devices belong to the same user or household\n* Probabilistically determine that two or more devices belong to the same user or household\n* Actively scan device characteristics for identification for probabilistic identification if users have allowed vendors to actively scan device characteristics for identification (Special Feature 2)"
		  },
		  "3": {
			"id": 3,
			"name": "Receive and use automatically-sent device characteristics for identification",
			"description": "Your device might be distinguished from other devices based on information it automatically sends, such as IP address or browser type.",
			"descriptionLegal": "Vendors can:\n* Create an identifier using data collected automatically from a device for specific characteristics, e.g. IP address, user-agent string.\n* Use such an identifier to attempt to re-identify a device.\nVendors cannot:\n* Create an identifier using data collected via actively scanning a device for specific characteristics, e.g. installed font or screen resolution without users’ separate opt-in to actively scanning device characteristics for identification.\n* Use such an identifier to re-identify a device."
		  }
		},
		"specialFeatures": {
		  "1": {
			"id": 1,
			"name": "Use precise geolocation data",
			"description": "Your precise geolocation data can be used in support of one or more purposes. This means your location can be accurate to within several meters.",
			"descriptionLegal": "Vendors can:\n* Collect and process precise geolocation data in support of one or more purposes.\nN.B. Precise geolocation means that there are no restrictions on the precision of a user’s location; this can be accurate to within several meters."
		  },
		  "2": {
			"id": 2,
			"name": "Actively scan device characteristics for identification",
			"description": "Your device can be identified based on a scan of your device's unique combination of characteristics.",
			"descriptionLegal": "Vendors can:\n* Create an identifier using data collected via actively scanning a device for specific characteristics, e.g. installed fonts or screen resolution.\n* Use such an identifier to re-identify a device."
		  }
		},
		"stacks": {
		  "1": {
			"id": 1,
			"purposes": [],
			"specialFeatures": [
			  1,
			  2
			],
			"name": "Precise geolocation data, and identification through device scanning",
			"description": "Precise geolocation and information about device characteristics can be used."
		  },
		  "2": {
			"id": 2,
			"purposes": [
			  2,
			  7
			],
			"specialFeatures": [],
			"name": "Basic ads, and ad measurement",
			"description": "Basic ads can be served. Ad performance can be measured."
		  },
		  "3": {
			"id": 3,
			"purposes": [
			  2,
			  3,
			  4
			],
			"specialFeatures": [],
			"name": "Personalised ads",
			"description": "Ads can be personalised based on a profile. More data can be added to better personalise ads."
		  },
		  "4": {
			"id": 4,
			"purposes": [
			  2,
			  7,
			  9
			],
			"specialFeatures": [],
			"name": "Basic ads, and ad measurement",
			"description": "Basic ads can be served. Ad performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "5": {
			"id": 5,
			"purposes": [
			  2,
			  3,
			  7
			],
			"specialFeatures": [],
			"name": "Basic ads, personalised ads profile, and ad measurement",
			"description": "Basic ads can be served. More data can be added to better personalise ads. Ad performance can be measured."
		  },
		  "6": {
			"id": 6,
			"purposes": [
			  2,
			  4,
			  7
			],
			"specialFeatures": [],
			"name": "Personalised ads display, and measurement",
			"description": "Ads can be personalised based on a profile. Ad performance can be measured."
		  },
		  "7": {
			"id": 7,
			"purposes": [
			  2,
			  4,
			  7,
			  9
			],
			"specialFeatures": [],
			"name": "Personalised ads display, ad measurement, and audience insights",
			"description": "Ads can be personalised based on a profile. Ad performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "8": {
			"id": 8,
			"purposes": [
			  2,
			  3,
			  4,
			  7
			],
			"specialFeatures": [],
			"name": "Personalised ads, and ad measurement",
			"description": "Ads can be personalised based on a profile. More data can be added to better personalise ads. Ad performance can be measured."
		  },
		  "9": {
			"id": 9,
			"purposes": [
			  2,
			  3,
			  4,
			  7,
			  9
			],
			"specialFeatures": [],
			"name": "Personalised ads, ad measurement, and audience insights",
			"description": "Ads can be personalised based on a profile. More data can be added to better personalise ads. Ad performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "10": {
			"id": 10,
			"purposes": [
			  3,
			  4
			],
			"specialFeatures": [],
			"name": "Personalised ads profile and display",
			"description": "Ads can be personalised based on a profile. More data can be added to better personalise ads."
		  },
		  "11": {
			"id": 11,
			"purposes": [
			  5,
			  6
			],
			"specialFeatures": [],
			"name": "Personalised content",
			"description": "Content can be personalised based on a profile. More data can be added to better personalise content."
		  },
		  "12": {
			"id": 12,
			"purposes": [
			  6,
			  8
			],
			"specialFeatures": [],
			"name": "Personalised content display, and content measurement",
			"description": "Content can be personalised based on a profile. Content performance can be measured."
		  },
		  "13": {
			"id": 13,
			"purposes": [
			  6,
			  8,
			  9
			],
			"specialFeatures": [],
			"name": "Personalised content display, content measurement and audience insights",
			"description": "Content can be personalised based on a profile. Content performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "14": {
			"id": 14,
			"purposes": [
			  5,
			  6,
			  8
			],
			"specialFeatures": [],
			"name": "Personalised content, and content measurement",
			"description": "Content can be personalised based on a profile. More data can be added to better personalise content. Content performance can be measured."
		  },
		  "15": {
			"id": 15,
			"purposes": [
			  5,
			  6,
			  8,
			  9
			],
			"specialFeatures": [],
			"name": "Personalised content, content measurement and audience insights",
			"description": "Content can be personalised based on a profile. More data can be added to better personalise content. Content performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "16": {
			"id": 16,
			"purposes": [
			  5,
			  6,
			  8,
			  9,
			  10
			],
			"specialFeatures": [],
			"name": "Personalised content, content measurement, audience insights, and product development",
			"description": "Content can be personalised based on a profile. More data can be added to better personalise content. Content performance can be measured. Insights about the audiences who saw the ads and content can be derived. Data can be used to build or improve user experience, systems, and software"
		  },
		  "17": {
			"id": 17,
			"purposes": [
			  7,
			  8,
			  9
			],
			"specialFeatures": [],
			"name": "Ad and content measurement, and audience insights",
			"description": "Ad and content performance can be measured.  Insights about the audiences who saw the ads and content can be derived."
		  },
		  "18": {
			"id": 18,
			"purposes": [
			  7,
			  8
			],
			"specialFeatures": [],
			"name": "Ad and content measurement",
			"description": "Ad and content performance can be measured."
		  },
		  "19": {
			"id": 19,
			"purposes": [
			  7,
			  9
			],
			"specialFeatures": [],
			"name": "Ad measurement, and audience insights",
			"description": "Ad can be measured.  Insights about the audiences who saw the ads and content can be derived."
		  },
		  "20": {
			"id": 20,
			"purposes": [
			  7,
			  8,
			  9,
			  10
			],
			"specialFeatures": [],
			"name": "Ad and content measurement, audience insights, and product development",
			"description": "Ad and content performance can be measured.  Insights about the audiences who saw the ads and content can be derived. Data can be used to build or improve user experience, systems, and software. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "21": {
			"id": 21,
			"purposes": [
			  8,
			  9,
			  10
			],
			"specialFeatures": [],
			"name": "Content measurement, audience insights, and product development.",
			"description": "Content performance can be measured. Insights about the audiences who saw the ads and content can be derived. Data can be used to build or improve user experience, systems, and software."
		  },
		  "22": {
			"id": 22,
			"purposes": [
			  8,
			  10
			],
			"specialFeatures": [],
			"name": "Content measurement, and product development",
			"description": "Content performance can be measured. Data can be used to build or improve user experience, systems, and software."
		  },
		  "23": {
			"id": 23,
			"purposes": [
			  2,
			  4,
			  6,
			  7,
			  8
			],
			"specialFeatures": [],
			"name": "Personalised ads and content display, ad and content measurement",
			"description": "Ads and content can be personalised based on a profile. Ad and content performance can be measured."
		  },
		  "24": {
			"id": 24,
			"purposes": [
			  2,
			  4,
			  6,
			  7,
			  8,
			  9
			],
			"specialFeatures": [],
			"name": "Personalised ads and content display, ad and content measurement, and audience insights",
			"description": "Ads and content can be personalised based on a profile. Ad and content performance can be measured. Insights about the audiences who saw the ads and content can be derived. Data can be used to build or improve user experience, systems, and software."
		  },
		  "25": {
			"id": 25,
			"purposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8
			],
			"specialFeatures": [],
			"name": "Personalised ads and content, ad and content measurement",
			"description": "Ads and content can be personalised based on a profile. More data can be added to better personalise ads and content. Ad and content performance can be measured."
		  },
		  "26": {
			"id": 26,
			"purposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9
			],
			"specialFeatures": [],
			"name": "Personalised ads and content, ad and content measurement, and audience insights",
			"description": "Ads and content can be personalised based on a profile. More data can be added to better personalise ads and content. Ad and content performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "27": {
			"id": 27,
			"purposes": [
			  3,
			  5
			],
			"specialFeatures": [],
			"name": "Personalised ads, and content profile",
			"description": "More data can be added to personalise ads and content."
		  },
		  "28": {
			"id": 28,
			"purposes": [
			  2,
			  4,
			  6
			],
			"specialFeatures": [],
			"name": "Personalised ads and content display",
			"description": "Ads and content can be personalised based on a profile."
		  },
		  "29": {
			"id": 29,
			"purposes": [
			  2,
			  7,
			  8,
			  9
			],
			"specialFeatures": [],
			"name": "Basic ads, ad and content measurement, and audience insights",
			"description": "Basic ads can be served. Ad and content performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "30": {
			"id": 30,
			"purposes": [
			  2,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9
			],
			"specialFeatures": [],
			"name": "Personalised ads display, personalised content, ad and content measurement, and audience insights",
			"description": "Ads and content can be personalised based on a profile. More data can be added to better personalise content. Ad and content performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "31": {
			"id": 31,
			"purposes": [
			  2,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialFeatures": [],
			"name": "Personalised ads display, personalised content, ad and content measurement, audience insights, and product development",
			"description": "Ads and content can be personalised based on a profile. More data can be added to better personalise content. Ad and content performance can be measured. Insights about the audiences who saw the ads and content can be derived. Data can be used to build or improve user experience, systems, and software."
		  },
		  "32": {
			"id": 32,
			"purposes": [
			  2,
			  5,
			  6,
			  7,
			  8,
			  9
			],
			"specialFeatures": [],
			"name": "Basic ads, personalised content, ad and content measurement, and audience insights",
			"description": "Basic ads can be served. Content can be personalised based on a profile. More data can be added to better personalise content. Ad and content performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "33": {
			"id": 33,
			"purposes": [
			  2,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialFeatures": [],
			"name": "Basic ads, personalised content, ad and content measurement, audience insights, and product development",
			"description": "Basic ads can be served. Content can be personalised based on a profile. More data can be added to better personalise content. Ad and content performance can be measured. Insights about the audiences who saw the ads and content can be derived. Data can be used to build or improve user experience, systems, and software."
		  },
		  "34": {
			"id": 34,
			"purposes": [
			  2,
			  5,
			  6,
			  8,
			  9
			],
			"specialFeatures": [],
			"name": "Basic ads, personalised content, content measurement, and audience insights",
			"description": "Basic ads can be served. Content can be personalised based on a profile. More data can be added to better personalise content. Ad and content performance can be measured. Insights about the audiences who saw the ads and content can be derived."
		  },
		  "35": {
			"id": 35,
			"purposes": [
			  2,
			  5,
			  6,
			  8,
			  9,
			  10
			],
			"specialFeatures": [],
			"name": "Basic ads, personalised content, content measurement, audience insights, and product development",
			"description": "Basic ads can be served. Content can be personalised based on a profile. More data can be added to better personalise content. Content performance can be measured. Insights about the audiences who saw the ads and content can be derived. Data can be used to build or improve user experience, systems, and software."
		  },
		  "36": {
			"id": 36,
			"purposes": [
			  2,
			  5,
			  6,
			  7
			],
			"specialFeatures": [],
			"name": "Basic ads, personalised content, and ad measurement",
			"description": "Basic ads can be served. Content can be personalised based on a profile. More data can be added to better personalise content. Ad performance can be measured."
		  },
		  "37": {
			"id": 37,
			"purposes": [
			  2,
			  5,
			  6,
			  7,
			  10
			],
			"specialFeatures": [],
			"name": "Basic ads, personalised content, ad measurement, and product development",
			"description": "Basic ads can be served. Content can be personalised based on a profile. More data can be added to better personalise content. Ad performance can be measured. Data can be used to build or improve user experience, systems, and software."
		  }
		},
		"vendors": {
		  "2": {
			"id": 2,
			"name": "Captify Technologies Limited",
			"purposes": [
			  1,
			  2,
			  3,
			  4
			],
			"legIntPurposes": [
			  7,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2
			],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "http://www.captify.co.uk/privacy-policy/"
		  },
		  "6": {
			"id": 6,
			"name": "AdSpirit GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  3
			],
			"specialFeatures": [],
			"policyUrl": "http://www.adspirit.de/privacy",
			"overflow": {
			  "httpGetLimit": 32
			}
		  },
		  "8": {
			"id": 8,
			"name": "Emerse Sverige AB",
			"purposes": [
			  1,
			  3,
			  4
			],
			"legIntPurposes": [
			  2,
			  7,
			  8,
			  9
			],
			"flexiblePurposes": [
			  2,
			  9
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www.emerse.com/privacy-policy/"
		  },
		  "9": {
			"id": 9,
			"name": "AdMaxim Inc.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://www.admaxim.com/admaxim-privacy-policy/"
		  },
		  "12": {
			"id": 12,
			"name": "BeeswaxIO Corporation",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  2
			],
			"features": [
			  1,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.beeswax.com/privacy/"
		  },
		  "15": {
			"id": 15,
			"name": "Adikteev / Emoteev",
			"purposes": [
			  1,
			  3,
			  4,
			  5,
			  6,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [
			  2,
			  7
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www.adikteev.com/privacy-policy-eng/"
		  },
		  "18": {
			"id": 18,
			"name": "Widespace AB",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7
			],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.widespace.com/legal/privacy-policy-notice/"
		  },
		  "23": {
			"id": 23,
			"name": "Amobee, Inc. ",
			"purposes": [
			  1,
			  2,
			  3,
			  4
			],
			"legIntPurposes": [
			  7,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.amobee.com/trust/privacy-guidelines"
		  },
		  "33": {
			"id": 33,
			"name": "ShareThis, Inc",
			"purposes": [
			  1,
			  3,
			  5,
			  6
			],
			"legIntPurposes": [
			  9,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://sharethis.com/privacy/"
		  },
		  "36": {
			"id": 36,
			"name": "RhythmOne LLC",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.rhythmone.com/privacy-policy"
		  },
		  "37": {
			"id": 37,
			"name": "NEURAL.ONE",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  7,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://web.neural.one/privacy-policy/"
		  },
		  "42": {
			"id": 42,
			"name": "Taboola Europe Limited",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.taboola.com/privacy-policy"
		  },
		  "44": {
			"id": 44,
			"name": "The ADEX GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://theadex.com/privacy-opt-out/"
		  },
		  "45": {
			"id": 45,
			"name": "Smart Adserver",
			"purposes": [
			  1,
			  4
			],
			"legIntPurposes": [
			  2,
			  7
			],
			"flexiblePurposes": [
			  2,
			  7
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://smartadserver.com/end-user-privacy-policy/"
		  },
		  "47": {
			"id": 47,
			"name": "ADMAN - Phaistos Networks, S.A.",
			"purposes": [
			  1,
			  2,
			  4,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://www.adman.gr/privacy"
		  },
		  "53": {
			"id": 53,
			"name": "Sirdata",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  2
			],
			"policyUrl": "https://www.sirdata.com/privacy/"
		  },
		  "58": {
			"id": 58,
			"name": "33Across",
			"purposes": [
			  1,
			  2,
			  7,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "http://www.33across.com/privacy-policy"
		  },
		  "59": {
			"id": 59,
			"name": "Sift Media, Inc",
			"purposes": [],
			"legIntPurposes": [
			  2
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.sift.co/privacy"
		  },
		  "61": {
			"id": 61,
			"name": "GumGum, Inc.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://gumgum.com/privacy-policy"
		  },
		  "65": {
			"id": 65,
			"name": "Location Sciences AI Ltd",
			"purposes": [
			  1,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.locationsciences.ai/privacy-policy/"
		  },
		  "66": {
			"id": 66,
			"name": "adsquare GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.adsquare.com/privacy"
		  },
		  "68": {
			"id": 68,
			"name": "Sizmek",
			"purposes": [
			  1,
			  3,
			  4
			],
			"legIntPurposes": [
			  2,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www.sizmek.com/privacy-policy/"
		  },
		  "72": {
			"id": 72,
			"name": "Nano Interactive GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://www.nanointeractive.com/privacy"
		  },
		  "78": {
			"id": 78,
			"name": "Flashtalking, Inc.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "http://www.flashtalking.com/privacypolicy/"
		  },
		  "80": {
			"id": 80,
			"name": "Sharethrough, Inc",
			"purposes": [
			  1,
			  2,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  4,
			  7,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://platform-cdn.sharethrough.com/privacy-policy"
		  },
		  "88": {
			"id": 88,
			"name": "TreSensa, Inc.",
			"purposes": [
			  1,
			  3,
			  4
			],
			"legIntPurposes": [
			  2,
			  7,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.tresensa.com/eu-privacy"
		  },
		  "89": {
			"id": 89,
			"name": "Tapad, Inc.",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www.tapad.com/eu-privacy-policy"
		  },
		  "90": {
			"id": 90,
			"name": "Teroa S.A.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.e-planning.net/en/privacy.html"
		  },
		  "93": {
			"id": 93,
			"name": "Adloox SA",
			"purposes": [],
			"legIntPurposes": [
			  7
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://adloox.com/disclaimer",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "94": {
			"id": 94,
			"name": "Blis Media Limited",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://www.blis.com/privacy/"
		  },
		  "100": {
			"id": 100,
			"name": "Fifty Technology Limited",
			"purposes": [
			  1,
			  2,
			  3,
			  4
			],
			"legIntPurposes": [
			  7,
			  9,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://fifty.io/privacy-policy.php"
		  },
		  "109": {
			"id": 109,
			"name": "LoopMe Limited",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://loopme.com/privacy-policy/"
		  },
		  "114": {
			"id": 114,
			"name": "Sublime",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9
			],
			"legIntPurposes": [
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://ayads.co/privacy.php"
		  },
		  "115": {
			"id": 115,
			"name": "smartclip Europe GmbH",
			"purposes": [
			  1,
			  3,
			  4
			],
			"legIntPurposes": [
			  2,
			  7,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://privacy-portal.smartclip.net/",
			"overflow": {
			  "httpGetLimit": 32
			}
		  },
		  "126": {
			"id": 126,
			"name": "DoubleVerify Inc.​",
			"purposes": [
			  2,
			  7,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.doubleverify.com/privacy/"
		  },
		  "127": {
			"id": 127,
			"name": "PIXIMEDIA SAS",
			"purposes": [
			  1,
			  3,
			  4,
			  5,
			  6,
			  9,
			  10
			],
			"legIntPurposes": [
			  2,
			  7,
			  8
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://piximedia.com/privacy/",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "128": {
			"id": 128,
			"name": "BIDSWITCH GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://www.bidswitch.com/privacy-policy/"
		  },
		  "129": {
			"id": 129,
			"name": "IPONWEB GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.iponweb.com/privacy-policy/"
		  },
		  "130": {
			"id": 130,
			"name": "NextRoll, Inc.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  2
			],
			"policyUrl": "https://www.nextroll.com/privacy"
		  },
		  "133": {
			"id": 133,
			"name": "digitalAudience",
			"purposes": [
			  1,
			  3,
			  5,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://digitalaudience.io/legal/privacy-cookies/"
		  },
		  "136": {
			"id": 136,
			"name": "Ströer SSP GmbH",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.stroeer.de/fileadmin/de/Konvergenz_und_Konzepte/Daten_und_Technologien/Stroeer_SSP/Downloads/Datenschutz_Stroeer_SSP.pdf"
		  },
		  "137": {
			"id": 137,
			"name": "Ströer SSP GmbH",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  9
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.stroeer.de/fileadmin/de/Konvergenz_und_Konzepte/Daten_und_Technologien/Stroeer_SSP/Downloads/Datenschutz_Stroeer_SSP.pdf"
		  },
		  "144": {
			"id": 144,
			"name": "district m inc.",
			"purposes": [
			  1,
			  2,
			  4,
			  7,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  2
			],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://districtm.net/en/page/platforms-data-and-privacy-policy/",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "147": {
			"id": 147,
			"name": "Adacado Technologies Inc. (DBA Adacado)",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.adacado.com/privacy-policy-april-25-2018/"
		  },
		  "153": {
			"id": 153,
			"name": "MADVERTISE MEDIA",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  5,
			  6,
			  10
			],
			"specialPurposes": [
			  2
			],
			"features": [],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://madvertise.com/en/gdpr/",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "163": {
			"id": 163,
			"name": "Bombora Inc.",
			"purposes": [
			  1,
			  3,
			  8
			],
			"legIntPurposes": [
			  7,
			  10
			],
			"flexiblePurposes": [
			  7,
			  10
			],
			"specialPurposes": [],
			"features": [
			  1
			],
			"specialFeatures": [],
			"policyUrl": "https://bombora.com/privacy"
		  },
		  "164": {
			"id": 164,
			"name": "Outbrain UK Ltd",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.outbrain.com/legal/privacy#privacy-policy"
		  },
		  "167": {
			"id": 167,
			"name": "Audiens S.r.l.",
			"purposes": [
			  1,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "http://www.audiens.com/privacy"
		  },
		  "174": {
			"id": 174,
			"name": "A Million Ads Ltd",
			"purposes": [],
			"legIntPurposes": [
			  2,
			  4
			],
			"flexiblePurposes": [
			  2,
			  4
			],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.amillionads.com/privacy-policy"
		  },
		  "177": {
			"id": 177,
			"name": "plista GmbH",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  7,
			  8,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.plista.com/about/privacy/"
		  },
		  "192": {
			"id": 192,
			"name": "remerge GmbH",
			"purposes": [
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://remerge.io/privacy-policy.html"
		  },
		  "199": {
			"id": 199,
			"name": "ADUX",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://www.adux.com/donnees-personelles/"
		  },
		  "203": {
			"id": 203,
			"name": "Revcontent, LLC",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6
			],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://intercom.help/revcontent2/en/articles/2290675-revcontent-s-privacy-policy"
		  },
		  "205": {
			"id": 205,
			"name": "Adssets AB",
			"purposes": [
			  2,
			  4
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://adssets.com/policy/"
		  },
		  "210": {
			"id": 210,
			"name": "Zemanta, Inc.",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  3,
			  7,
			  9,
			  10
			],
			"flexiblePurposes": [
			  3,
			  7,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1
			],
			"specialFeatures": [],
			"policyUrl": "http://www.zemanta.com/legal/privacy"
		  },
		  "213": {
			"id": 213,
			"name": "emetriq GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  2
			],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www.emetriq.com/datenschutz/"
		  },
		  "215": {
			"id": 215,
			"name": "ARMIS SAS",
			"purposes": [
			  1,
			  2,
			  7
			],
			"legIntPurposes": [
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://armis.tech/en/armis-personal-data-privacy-policy/"
		  },
		  "218": {
			"id": 218,
			"name": "Bigabid Media ltd",
			"purposes": [],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.bigabid.com/privacy-policy"
		  },
		  "224": {
			"id": 224,
			"name": "adrule mobile GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  8
			],
			"legIntPurposes": [
			  7
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7,
			  8
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.adrule.net/de/datenschutz/"
		  },
		  "226": {
			"id": 226,
			"name": "Publicis Media GmbH",
			"purposes": [
			  1,
			  9
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  10
			],
			"specialPurposes": [],
			"features": [
			  1
			],
			"specialFeatures": [],
			"policyUrl": "https://www.publicismedia.de/datenschutz/"
		  },
		  "228": {
			"id": 228,
			"name": "McCann Discipline LTD",
			"purposes": [
			  1,
			  2,
			  5,
			  6,
			  7,
			  8
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.primis.tech/wp-content/uploads/2020/01/Primis-Privacy-Policy.pdf"
		  },
		  "235": {
			"id": 235,
			"name": "Bucksense Inc",
			"purposes": [
			  2,
			  3,
			  4,
			  7,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7,
			  9
			],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://www.bucksense.com/platform-privacy-policy/"
		  },
		  "240": {
			"id": 240,
			"name": "7Hops.com Inc. (ZergNet)",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  5,
			  6,
			  8,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://zergnet.com/privacy"
		  },
		  "243": {
			"id": 243,
			"name": "Cloud Technologies S.A.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.cloudtechnologies.pl/en/internet-advertising-privacy-policy"
		  },
		  "248": {
			"id": 248,
			"name": "Converge-Digital",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2
			],
			"flexiblePurposes": [
			  2
			],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://converge-digital.com/privacy-policy/",
			"overflow": {
			  "httpGetLimit": 32
			}
		  },
		  "252": {
			"id": 252,
			"name": "Jaduda GmbH",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.jadudamobile.com/datenschutzerklaerung/"
		  },
		  "256": {
			"id": 256,
			"name": "Bounce Exchange, Inc",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  7,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.bouncex.com/privacy/"
		  },
		  "262": {
			"id": 262,
			"name": "Fyber ",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.fyber.com/legal/privacy-policy/"
		  },
		  "272": {
			"id": 272,
			"name": "A.Mob",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.we-are-adot.com/privacy-policy/"
		  },
		  "273": {
			"id": 273,
			"name": "Bannerflow AB",
			"purposes": [
			  1,
			  4
			],
			"legIntPurposes": [
			  7
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.bannerflow.com/privacy "
		  },
		  "277": {
			"id": 277,
			"name": "Codewise VL Sp. z o.o. Sp. k",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9
			],
			"legIntPurposes": [
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://voluumdsp.com/end-user-privacy-policy/"
		  },
		  "278": {
			"id": 278,
			"name": "Integral Ad Science, Inc.",
			"purposes": [],
			"legIntPurposes": [
			  7,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://integralads.com/privacy-policy/"
		  },
		  "279": {
			"id": 279,
			"name": "Mirando GmbH &amp; Co KG",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  7
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://wwwmirando.de/datenschutz/"
		  },
		  "281": {
			"id": 281,
			"name": "Wizaly",
			"purposes": [
			  1,
			  7,
			  8,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www.wizaly.com/terms-of-use#privacy-policy"
		  },
		  "285": {
			"id": 285,
			"name": "Comcast International France SAS",
			"purposes": [
			  1
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.freewheel.com/privacy-policy"
		  },
		  "294": {
			"id": 294,
			"name": "Jivox Corporation",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.jivox.com/privacy"
		  },
		  "302": {
			"id": 302,
			"name": "Mobile Professionals BV",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  7,
			  8,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  3,
			  5,
			  7,
			  8
			],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://mobpro.com/privacy.html"
		  },
		  "303": {
			"id": 303,
			"name": "Orion Semantics",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://static.orion-semantics.com/privacy.html"
		  },
		  "304": {
			"id": 304,
			"name": "On Device Research Limited",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://s.on-device.com/privacyPolicy"
		  },
		  "314": {
			"id": 314,
			"name": "Keymantics",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.keymantics.com/assets/privacy-policy.pdf"
		  },
		  "318": {
			"id": 318,
			"name": "Accorp Sp. z o.o.",
			"purposes": [
			  1,
			  3,
			  4,
			  9
			],
			"legIntPurposes": [
			  7,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "http://www.instytut-pollster.pl/privacy-policy/"
		  },
		  "325": {
			"id": 325,
			"name": "Knorex Pte Ltd",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.knorex.com/privacy"
		  },
		  "328": {
			"id": 328,
			"name": "Gemius SA",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.gemius.com/cookie-policy.html"
		  },
		  "350": {
			"id": 350,
			"name": "Free Stream Media Corp. dba Samba TV",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://samba.tv/legal/privacy-policy/",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "351": {
			"id": 351,
			"name": "Samba TV UK Limited",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://samba.tv/legal/privacy-policy/",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "358": {
			"id": 358,
			"name": "MGID Inc.",
			"purposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.mgid.com/privacy-policy"
		  },
		  "371": {
			"id": 371,
			"name": "Seeding Alliance GmbH",
			"purposes": [],
			"legIntPurposes": [
			  2,
			  7,
			  8
			],
			"flexiblePurposes": [
			  7
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://seeding-alliance.de/datenschutz"
		  },
		  "374": {
			"id": 374,
			"name": "Bmind a Sales Maker Company, S.L.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [],
			"policyUrl": "http://www.bmind.es/legal-notice/"
		  },
		  "387": {
			"id": 387,
			"name": "Triapodi Ltd.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://appreciate.mobi/page.html#/end-user-privacy-policy"
		  },
		  "402": {
			"id": 402,
			"name": "Effiliation",
			"purposes": [
			  1,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://inter.effiliation.com/politique-confidentialite.html"
		  },
		  "413": {
			"id": 413,
			"name": "Eulerian Technologies",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.eulerian.com/en/privacy/"
		  },
		  "415": {
			"id": 415,
			"name": "Seenthis AB",
			"purposes": [],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://seenthis.co/privacy-notice-2018-04-18.pdf"
		  },
		  "422": {
			"id": 422,
			"name": "Brand Metrics Sweden AB",
			"purposes": [
			  1,
			  6,
			  7,
			  8,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://collector.brandmetrics.com/brandmetrics_privacypolicy.pdf"
		  },
		  "423": {
			"id": 423,
			"name": "travel audience GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://travelaudience.com/product-privacy-policy/"
		  },
		  "424": {
			"id": 424,
			"name": "KUPONA GmbH",
			"purposes": [
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.kupona.de/dsgvo/"
		  },
		  "428": {
			"id": 428,
			"name": "Internet BillBoard a.s.",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  7
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://www.ibillboard.com/en/privacy-information/"
		  },
		  "429": {
			"id": 429,
			"name": "Signals",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://signalsdata.com/platform-cookie-policy/"
		  },
		  "436": {
			"id": 436,
			"name": "INVIBES GROUP",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  9,
			  10
			],
			"legIntPurposes": [
			  7,
			  8
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "http://www.invibes.com/terms"
		  },
		  "439": {
			"id": 439,
			"name": "Bit Q Holdings Limited",
			"purposes": [
			  1,
			  7,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.rippll.com/privacy"
		  },
		  "440": {
			"id": 440,
			"name": "DEFINE MEDIA GMBH",
			"purposes": [
			  1,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [
			  2
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://www.definemedia.de/datenschutz-conative/"
		  },
		  "447": {
			"id": 447,
			"name": "Adludio Ltd",
			"purposes": [
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.adludio.com/privacy-policy/"
		  },
		  "450": {
			"id": 450,
			"name": "Neodata Group srl",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.neodatagroup.com/en/security-policy",
			"overflow": {
			  "httpGetLimit": 32
			}
		  },
		  "466": {
			"id": 466,
			"name": "TACTIC™ Real-Time Marketing AS",
			"purposes": [],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://tacticrealtime.com/privacy/"
		  },
		  "467": {
			"id": 467,
			"name": "Haensel AMS GmbH",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  7
			],
			"flexiblePurposes": [
			  7
			],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://haensel-ams.com/data-privacy/"
		  },
		  "486": {
			"id": 486,
			"name": "Madington",
			"purposes": [],
			"legIntPurposes": [
			  7,
			  8,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://delivered-by-madington.com/dat-privacy-policy/"
		  },
		  "490": {
			"id": 490,
			"name": "PLAYGROUND XYZ EMEA LTD",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7
			],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://playground.xyz/privacy"
		  },
		  "491": {
			"id": 491,
			"name": "Triboo Data Analytics",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.shinystat.com/it/informativa_sito.html"
		  },
		  "495": {
			"id": 495,
			"name": "Arcspire Limited",
			"purposes": [
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://public.arcspire.io/privacy.pdf",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "498": {
			"id": 498,
			"name": "Dr. Banner",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8
			],
			"legIntPurposes": [
			  9,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://drbanner.com/privacypolicy_en/"
		  },
		  "502": {
			"id": 502,
			"name": "NEXD",
			"purposes": [
			  7,
			  10
			],
			"legIntPurposes": [
			  9
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://nexd.com/privacy-policy",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "507": {
			"id": 507,
			"name": "AdsWizz Inc.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.adswizz.com/our-privacy-policy/"
		  },
		  "512": {
			"id": 512,
			"name": "PubNative GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  8,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [
			  2
			],
			"policyUrl": "https://pubnative.net/privacy-notice/"
		  },
		  "516": {
			"id": 516,
			"name": "Pexi B.V.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://pexi.nl/privacy-policy/"
		  },
		  "522": {
			"id": 522,
			"name": "Tealium Inc.",
			"purposes": [
			  1,
			  3,
			  5,
			  8,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://tealium.com/privacy-policy/"
		  },
		  "524": {
			"id": 524,
			"name": "The Ozone Project Limited",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://ozoneproject.com/privacy-policy"
		  },
		  "530": {
			"id": 530,
			"name": "Near Pte Ltd",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://near.co/privacy",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "531": {
			"id": 531,
			"name": "Smartclip Hispania SL",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://rgpd-smartclip.com/"
		  },
		  "535": {
			"id": 535,
			"name": "INNITY",
			"purposes": [
			  1,
			  2,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www.innity.com/privacy-policy.php"
		  },
		  "543": {
			"id": 543,
			"name": "PaperG, Inc. dba Thunder Industries",
			"purposes": [
			  1,
			  3,
			  4,
			  5,
			  6
			],
			"legIntPurposes": [
			  2,
			  7,
			  8
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.makethunder.com/privacy"
		  },
		  "544": {
			"id": 544,
			"name": "Kochava Inc.",
			"purposes": [],
			"legIntPurposes": [
			  7
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.kochava.com/support-privacy/"
		  },
		  "545": {
			"id": 545,
			"name": "Reignn Platform Ltd",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  8,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [
			  2
			],
			"policyUrl": "http://reignn.com/user-privacy-policy"
		  },
		  "546": {
			"id": 546,
			"name": "Smart Traffik",
			"purposes": [
			  1,
			  8
			],
			"legIntPurposes": [
			  7
			],
			"flexiblePurposes": [
			  7,
			  8
			],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://smart-traffik.io/politique-de-confidentialite"
		  },
		  "549": {
			"id": 549,
			"name": "Bandsintown Amplified LLC",
			"purposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://corp.bandsintown.com/privacy"
		  },
		  "550": {
			"id": 550,
			"name": "Happydemics",
			"purposes": [
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.iubenda.com/privacy-policy/69056167/full-legal"
		  },
		  "553": {
			"id": 553,
			"name": "Adhese",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://adhese.com/privacy-and-cookie-policy"
		  },
		  "554": {
			"id": 554,
			"name": "RMSi Radio Marketing Service interactive GmbH",
			"purposes": [
			  1,
			  3,
			  4
			],
			"legIntPurposes": [
			  2,
			  7,
			  9
			],
			"flexiblePurposes": [
			  2
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.rms.de/datenschutz/"
		  },
		  "556": {
			"id": 556,
			"name": "adhood.com",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://v3.adhood.com/en/site/politikavekurallar/gizlilik.php?lang=en"
		  },
		  "559": {
			"id": 559,
			"name": "Otto (GmbH &amp; Co KG)",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [
			  2
			],
			"policyUrl": "https://www.otto.de/shoppages/service/datenschutz"
		  },
		  "568": {
			"id": 568,
			"name": "Jointag S.r.l.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.jointag.com/privacy/kariboo/publisher/third/"
		  },
		  "571": {
			"id": 571,
			"name": "ViewPay",
			"purposes": [
			  1,
			  2,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [],
			"policyUrl": "http://viewpay.tv/mentions-legales/"
		  },
		  "580": {
			"id": 580,
			"name": "Goldbach Group AG",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://goldbach.com/ch/de/datenschutz"
		  },
		  "587": {
			"id": 587,
			"name": "Localsensor B.V.",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.localsensor.com/privacy.html"
		  },
		  "593": {
			"id": 593,
			"name": "Programatica de publicidad S.L.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://datmean.com/politica-privacidad/"
		  },
		  "602": {
			"id": 602,
			"name": "Online Solution Int Limited",
			"purposes": [
			  2,
			  7,
			  8,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  7,
			  8,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://adsafety.net/privacy.html",
			"overflow": {
			  "httpGetLimit": 32
			}
		  },
		  "606": {
			"id": 606,
			"name": "Impactify ",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://impactify.io/privacy-policy/"
		  },
		  "607": {
			"id": 607,
			"name": "ucfunnel Co., Ltd.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.ucfunnel.com/privacy-policy"
		  },
		  "609": {
			"id": 609,
			"name": "Predicio",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "http://www.predic.io/privacy",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "612": {
			"id": 612,
			"name": "Adnami Aps",
			"purposes": [
			  1
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.adnami.io/privacy"
		  },
		  "613": {
			"id": 613,
			"name": "Adserve.zone / Artworx AS",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  7
			],
			"flexiblePurposes": [
			  2
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://adserve.zone/adserveprivacypolicy.html"
		  },
		  "618": {
			"id": 618,
			"name": "BEINTOO SPA",
			"purposes": [
			  1,
			  2,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://www.beintoo.com/privacy-cookie-policy/"
		  },
		  "620": {
			"id": 620,
			"name": "Blue",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://www.getblue.io/privacy/"
		  },
		  "626": {
			"id": 626,
			"name": "Hivestack Inc.",
			"purposes": [
			  1,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://hivestack.com/privacy-policy"
		  },
		  "628": {
			"id": 628,
			"name": ": Tappx",
			"purposes": [
			  1,
			  2,
			  4,
			  7,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.tappx.com/en/privacy-policy/"
		  },
		  "648": {
			"id": 648,
			"name": "TrueData Solutions, Inc.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.truedata.co/privacy-policy/"
		  },
		  "650": {
			"id": 650,
			"name": "Telefonica Investigación y Desarrollo S.A.U",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://www.cognitivemarketing.tid.es/"
		  },
		  "652": {
			"id": 652,
			"name": "Skaze",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://www.skaze.fr/rgpd/"
		  },
		  "653": {
			"id": 653,
			"name": "Smartme Analytics",
			"purposes": [
			  7,
			  8,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "http://smartmeapp.com/info/smartme/aviso_legal.php"
		  },
		  "656": {
			"id": 656,
			"name": "Think Clever Media",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  9
			],
			"legIntPurposes": [
			  7,
			  8,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.contentignite.com/privacy-policy/"
		  },
		  "657": {
			"id": 657,
			"name": "GP One GmbH",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.gsi-one.org/de/privacy-policy.html"
		  },
		  "659": {
			"id": 659,
			"name": "Research and Analysis of Media in Sweden AB",
			"purposes": [],
			"legIntPurposes": [
			  7,
			  8,
			  9
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www2.rampanel.com/privacy-policy/"
		  },
		  "662": {
			"id": 662,
			"name": "Soundcast",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://soundcast.fm/en/data-privacy"
		  },
		  "663": {
			"id": 663,
			"name": "Mobsuccess",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  7
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1
			],
			"specialFeatures": [],
			"policyUrl": "https://www.mobsuccess.com/en/privacy"
		  },
		  "665": {
			"id": 665,
			"name": "Digital East GmbH",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.digitaleast.mobi/en/legal/privacy-policy/"
		  },
		  "668": {
			"id": 668,
			"name": "WhatRocks Inc. ",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.whatrocks.co/en/privacy-policy "
		  },
		  "674": {
			"id": 674,
			"name": "Duration Media, LLC.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.durationmedia.net/privacy-policy"
		  },
		  "675": {
			"id": 675,
			"name": "Instreamatic inc.",
			"purposes": [
			  2,
			  3,
			  4,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://instreamatic.com/privacy-policy/"
		  },
		  "676": {
			"id": 676,
			"name": "BusinessClick",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.businessclick.com/documents/RegulaminProgramuBusinessClick-2019.pdf"
		  },
		  "681": {
			"id": 681,
			"name": "MyTraffic",
			"purposes": [
			  1,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  9
			],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.mytraffic.io/en/privacy"
		  },
		  "683": {
			"id": 683,
			"name": "Cookie Market LTD",
			"purposes": [
			  1,
			  3,
			  4,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "http://cookie.market/privacyPolicy.php"
		  },
		  "684": {
			"id": 684,
			"name": "Blue Billywig BV",
			"purposes": [],
			"legIntPurposes": [
			  7
			],
			"flexiblePurposes": [
			  7
			],
			"specialPurposes": [
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.bluebillywig.com/privacy-statement/"
		  },
		  "686": {
			"id": 686,
			"name": "The MediaGrid Inc.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www.themediagrid.com/privacy-policy/"
		  },
		  "687": {
			"id": 687,
			"name": "MISSENA",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2
			],
			"specialPurposes": [
			  2
			],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "http://missena.com/confidentialite/"
		  },
		  "688": {
			"id": 688,
			"name": "Effinity",
			"purposes": [],
			"legIntPurposes": [
			  2
			],
			"flexiblePurposes": [
			  2,
			  7
			],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.effiliation.com/politique-de-confidentialite/",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "690": {
			"id": 690,
			"name": "Go.pl sp. z o.o.",
			"purposes": [
			  1,
			  2,
			  3,
			  4
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://go.pl/polityka-prywatnosci/"
		  },
		  "691": {
			"id": 691,
			"name": "Lifesight Pte. Ltd.",
			"purposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.lifesight.io/privacy-policy/"
		  },
		  "694": {
			"id": 694,
			"name": "Snapupp Technologies SL",
			"purposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  9
			],
			"legIntPurposes": [
			  7,
			  8,
			  10
			],
			"flexiblePurposes": [
			  10
			],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.enterprise.noddus.com/privacy-policy",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "699": {
			"id": 699,
			"name": "HyperTV Inc.",
			"purposes": [
			  1,
			  2,
			  4,
			  9
			],
			"legIntPurposes": [
			  7,
			  10
			],
			"flexiblePurposes": [],
			"specialPurposes": [
			  2
			],
			"features": [
			  1
			],
			"specialFeatures": [],
			"policyUrl": "https://www.hypertvx.com/privacy/"
		  },
		  "702": {
			"id": 702,
			"name": "Kwanko",
			"purposes": [
			  1,
			  2,
			  7,
			  8
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  7,
			  8
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.kwanko.com/fr/rgpd/"
		  },
		  "703": {
			"id": 703,
			"name": "MindTake Research GmbH",
			"purposes": [
			  1,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1
			],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://www.mindtake.com/en/reppublika-privacy-policy"
		  },
		  "707": {
			"id": 707,
			"name": "Dentsu Aegis Network Italia SpA",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.dentsuaegisnetwork.com/it/it/policies/info-cookie"
		  },
		  "708": {
			"id": 708,
			"name": "Dugout Limited ",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://dugout.com/privacy-policy"
		  },
		  "709": {
			"id": 709,
			"name": "NC Audience Exchange, LLC (NewsIQ)",
			"purposes": [
			  1,
			  3,
			  4,
			  5
			],
			"legIntPurposes": [
			  2,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "www.ncaudienceexchange.com/privacy"
		  },
		  "711": {
			"id": 711,
			"name": "SITU8ED SA",
			"purposes": [
			  1,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.situ8ed.com/privacy-policy/"
		  },
		  "712": {
			"id": 712,
			"name": "Inspired Mobile Limited",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://byinspired.com/privacypolicy.pdf"
		  },
		  "713": {
			"id": 713,
			"name": "Dataseat Ltd",
			"purposes": [
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  7,
			  8
			],
			"specialPurposes": [],
			"features": [
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://dataseat.com/privacy-policy"
		  },
		  "714": {
			"id": 714,
			"name": "Survata Inc.",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  7,
			  9
			],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.survata.com/respondent-privacy-policy/",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "716": {
			"id": 716,
			"name": "OnAudience Ltd",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.onaudience.com/internet-advertising-privacy-policy"
		  },
		  "719": {
			"id": 719,
			"name": "Online Advertising Network Sp. z o.o.",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://www.oan.pl/en/privacy-policy"
		  },
		  "720": {
			"id": 720,
			"name": "AAX LLC",
			"purposes": [
			  1,
			  3,
			  4
			],
			"legIntPurposes": [
			  2,
			  7
			],
			"flexiblePurposes": [
			  3,
			  4,
			  7
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://aax.media/privacy/",
			"overflow": {
			  "httpGetLimit": 32
			}
		  },
		  "722": {
			"id": 722,
			"name": "agof services gmbh",
			"purposes": [],
			"legIntPurposes": [
			  3,
			  5,
			  7,
			  8,
			  9,
			  10
			],
			"flexiblePurposes": [
			  5,
			  8,
			  9,
			  10
			],
			"specialPurposes": [
			  1
			],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [
			  2
			],
			"policyUrl": "http://www.agof.de/datenschutz/"
		  },
		  "723": {
			"id": 723,
			"name": "Adzymic Pte Ltd",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  8
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "http://www.adzymic.co/privacy",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "725": {
			"id": 725,
			"name": "Pubfinity LLC",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://pubfinity.com/privacy-policy/",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "726": {
			"id": 726,
			"name": "YouGov",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  7,
			  8,
			  9,
			  10
			],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://yougov.co.uk/about/terms-combined/#/privacy"
		  },
		  "729": {
			"id": 729,
			"name": "Cavai AS & UK ",
			"purposes": [],
			"legIntPurposes": [
			  7,
			  8
			],
			"flexiblePurposes": [
			  7,
			  8
			],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://cav.ai/privacy-policy/"
		  },
		  "731": {
			"id": 731,
			"name": "GeistM Technologies LTD",
			"purposes": [],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://www.geistm.com/privacy"
		  },
		  "733": {
			"id": 733,
			"name": "Anzu Virtual reality LTD",
			"purposes": [
			  1,
			  2,
			  4,
			  7,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1,
			  2
			],
			"specialFeatures": [
			  1
			],
			"policyUrl": "https://anzu.io/privacy/"
		  },
		  "734": {
			"id": 734,
			"name": "Cint AB",
			"purposes": [
			  7,
			  9
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2
			],
			"specialFeatures": [],
			"policyUrl": "https://www.cint.com/participant-privacy-notice/",
			"overflow": {
			  "httpGetLimit": 32
			}
		  },
		  "735": {
			"id": 735,
			"name": "Deutsche Post AG",
			"purposes": [
			  1,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  1
			],
			"specialFeatures": [],
			"policyUrl": "https://www.deutschepost.de/de/f/footer/datenschutz.html"
		  },
		  "737": {
			"id": 737,
			"name": "Monet Engine Inc",
			"purposes": [
			  1,
			  2,
			  7
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [],
			"specialFeatures": [],
			"policyUrl": "https://appmonet.com/privacy-policy/"
		  },
		  "740": {
			"id": 740,
			"name": "6Sense Insights, Inc.",
			"purposes": [
			  1
			],
			"legIntPurposes": [
			  2,
			  3,
			  4,
			  5,
			  7,
			  8,
			  10
			],
			"flexiblePurposes": [
			  2,
			  3,
			  4,
			  5,
			  7,
			  8,
			  10
			],
			"specialPurposes": [],
			"features": [
			  1,
			  2,
			  3
			],
			"specialFeatures": [],
			"policyUrl": "https://6sense.com/privacy-policy/",
			"overflow": {
			  "httpGetLimit": 128
			}
		  },
		  "741": {
			"id": 741,
			"name": "Brand Advance Limited",
			"purposes": [
			  1,
			  2,
			  3,
			  4,
			  5,
			  6,
			  7,
			  8,
			  9,
			  10
			],
			"legIntPurposes": [],
			"flexiblePurposes": [],
			"specialPurposes": [],
			"features": [
			  2,
			  3
			],
			"specialFeatures": [
			  1,
			  2
			],
			"policyUrl": "https://www.wearebrandadvance.com/website-privacy-policy"
		  },
		  "744": {
			"id": 744,
			"name": "Vidazoo Ltd",
			"purposes": [
			  3,
			  4
			],
			"legIntPurposes": [
			  2,
			  7,
			  8,
			  10
			],
			"flexiblePurposes": [
			  3,
			  4
			],
			"specialPurposes": [
			  1,
			  2
			],
			"features": [
			  2
			],
			"specialFeatures": [
			  2
			],
			"policyUrl": "https://vidazoo.gitbook.io/vidazoo-legal/privacy-policy",
			"overflow": {
			  "httpGetLimit": 128
			}
		  }
		}
	  }`
}
