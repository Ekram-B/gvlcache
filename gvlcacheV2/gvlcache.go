package gvlcachev2

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ezoic/ezcache"
	"github.com/ezoic/publisher-backend/utils"
)

// GVLVersionTwoValue is built to match the formatting of the GVL Version 2 based on the
// specification set by the IAB lab https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/TCFv2/IAB%20Tech%20Lab%20-%20Consent%20string%20and%20vendor%20list%20formats%20v2.md#caching-the-global-vendor-list
type GVLVersionTwoValue struct {
	GVLSpecificationVersion string `json:"gvlSpecificationVersion"`
	VendorListVersion       string `json:"vendorListVersion"`
	TCFPolicyVersion        string `json:"tcfPolicyVersion "`
	LastUpdated             string `json:"lastUpdated"`
	// map of integer to purposes
	Purposes map[int]GVLVersionTwoPurpose `json:"purposes"`
	// map of integer to special purposes
	SpecialPurposes map[int]GVLVersionTwoSpecialPurpose `json:"specialPurposes"`
	// map of integer to features
	Features map[int]GVLVersionTwoPurpose `json:"features"`
	// map of intege to special features
	SpecialFeatures map[int]GVLVersionTwoSpecialFeature `json:"specialFeatures"`
	// map of integer to vendors
	Vendors map[int]GVLVersionTwoVendor `json:"vendors"`
	// map of integer to stacks
	Stacks map[int]GVLVersionTwoStack `json:"stacks"`
}

// GVLVersionTwoPurpose can have ids between id=0 and up to no higher than id=24 inclusive on both sides
type GVLVersionTwoPurpose struct {
	ID               int    `json:"id"`               //REQUIRED
	Name             string `json:"name"`             //REQUIRED
	Description      string `json:"description"`      //REQUIRED
	DescriptionLegal string `json:"descriptionLegal"` //REQUIRED
	Consentable      bool   `json:"consentable"`      //OPTIONAL, default=true  false means CMPs should never afford users the means to provide an opt-in consent choice
	RightToObject    bool   `json:"rightToObject"`    //OPTIONAL, default=true  false means CMPs should never afford users the means to exercise a right to object
}

type GVLVersionTwoSpecialPurpose struct {
	ID               int    `json:"id"`               //REQUIRED
	Name             string `json:"name"`             //REQUIRED
	Description      string `json:"description"`      //REQUIRED
	DescriptionLegal string `json:"descriptionLegal"` //REQUIRED
	Consentable      bool   `json:"consentable"`      //OPTIONAL, default=true  false means CMPs should never afford users the means to provide an opt-in consent choice
	RightToObject    bool   `json:"rightToObject"`    //OPTIONAL, default=true  false means CMPs should never afford users the means to exercise a right to object
}

// GVLVersionTwoFeature can have ids between id=0 to no higher than id=64 inclusive on both sides
type GVLVersionTwoFeature struct {
	ID               int    `json:"id"`               //REQUIRED
	Name             string `json:"name"`             //REQUIRED
	Description      string `json:"description"`      //REQUIRED
	DescriptionLegal string `json:"descriptionLegal"` //REQUIRED
	Consentable      bool   `json:"consentable"`      //OPTIONAL, default=true  false means CMPs should never afford users the means to provide an opt-in consent choice
	RightToObject    bool   `json:"rightToObject"`    //OPTIONAL, default=true  false means CMPs should never afford users the means to exercise a right to object
}

// Special features differ from simple features in that CMPs MUST provide
// users with a means to signal an opt-in choice as to whether vendors
// may employ the feature when performing any purpose processing.

// GVLVersionTwoSpecialFeature can have ids between id=0 to no higher than id=8 inclusive on both sides
type GVLVersionTwoSpecialFeature struct {
	ID               int    `json:"id"`               //REQUIRED
	Name             string `json:"name"`             //REQUIRED
	Description      string `json:"description"`      //REQUIRED
	DescriptionLegal string `json:"descriptionLegal"` //REQUIRED
	Consentable      bool   `json:"consentable"`      //OPTIONAL, default=true  false means CMPs should never afford users the means to provide an opt-in consent choice
	RightToObject    bool   `json:"rightToObject"`    //OPTIONAL, default=true  false means CMPs should never afford users the means to exercise a right to object
}

/*
- Constraints on vendor object
		1.  Either purposes OR legIntPurposes can be missing/empty, but not both.
		2.  A Purpose id must not be present in both purposes and legIntPurposes
		3.  A Purpose id listed in flexiblePurposes must have been declared in one of purposes or legIntPurposes.
		4.  Purpose id values included in the three purpose fields must be in the range from 1 to N, where N is the highest purpose id published in this GVL file.
		5.  "features": array of positive integers, OPTIONAL. Array may be empty. List of Features the Vendor may utilize when performing some declared Purposes processing.
		6.  "specialFeatures": array of positive integers, OPTIONAL. Array may be empty. List of Special Features the Vendor may utilize when performing some declared Purposes processing.
		7.  "SpecialPurposes": array of positive integers, OPTIONAL. Array may be empty. List of Special Purposes declared as performed on the legal basis of a legitimate interest
		8.  "policyUrl": url string, REQUIRED URL to the Vendor's privacy policy document.
		9.  "deletedDate": date string ("2019-05-28T00:00:00Z") OPTIONAL, If present, vendor is considered deleted after this date/time and MUST NOT be established to users.
		10. "overflow": object specifying the vendor's http GET request length  limit OPTIONAL. Has the following members & values
*/

type GVLVersionTwoVendor struct {
	ID               int                   `json:"id"`               //REQUIRED
	Name             string                `json:"name"`             //REQUIRED
	Purposes         []int                 `json:"purposes"`         //!!!!REQUIRED!!!!! (refer to constraints)
	SpecialPurposes  []int                 `json:"specialPurposes"`  //OPTIONAL
	LegIntPurposes   []int                 `json:"legIntPurposes"`   //!!!!REQUIRED!!!!! (refer to constraints)The array may be empty. The array must be of exclusively positive integers.
	FlexiblePurposes []int                 `json:"flexiblePurposes"` //OPTIONAL
	Features         []int                 `json:"features"`         //OPTIONAL
	SpecialFeatures  []int                 `json:"specialFeatures"`  //OPTIONAL
	PolicyURL        string                `json:"policyUrl"`        //REQUIRED
	DeletedDate      string                `json:"deletedDate"`      //OPTIONAL
	Overflow         GVLVersionTwoOverflow `json:"overflow"`         //OPTIONAL
}

type GVLVersionTwoOverflow struct {
	//  32 or 128 are supported options
	httpGetLimit int `json:"httpGetLimit"`
}
type GVLVersionTwoStack struct {
	ID              int    `json:"id"`
	Purposes        []int  `json:"purposes"`
	SpecialFeatures []int  `json:"specialFeatures"`
	Name            string `json:"name"`
	Description     string `json:"description"`
}

func (gvl *GVLVersionTwoValue) getGVLVersionTwoValueFromCache() bool {

	// err := ezcache.LoadKeyObject("middleton", gvlCacheKey, *gvl)
	// if err == nil {
	// 	log.Print("Successfully loaded GVL value from the cache!")
	// 	return true
	// }
	//  There was an error returned from trying to retreive the cached value -
	//  possibly b/c there wasn't a cached value that already existed
	return false
}

func (gvl *GVLVersionTwoValue) isIABResponseBodyMalformed(body []byte) bool {
	// Use what we know about what is suppose to be within the response, as well as the constraints that are
	// suppose to be met in order to implement this

	return true
}

func (gvl *GVLVersionTwoValue) getGVLVersionTwoValueFromIABSource() error {
	var url string

	if utils.IsLocal() {
		// url for locally set up dummy server
		url = "http://127.0.0.1/8085/v2/vendor-list.json"
	} else {
		// url for the !!actual IAB server!! (for now this is set as local host in order to avoid issues with calling the IAB server)
		// this can only be done once per caching time! Otherwise, we may be blocked from requesting it
		url = "http://127.0.0.1/8085/v2/vendor-list.json"
		// actual IAB Server URL -> https://vendorlist.consensu.org/v2/vendor-list.json
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	// There are some constraints to be met (assuming IAB provides the correct response content), that begin to matter at this point of the code.
	// 1. A correct GVL shall be one that matches the most recently updated and publicly available version JSON  retreivable from the link provided by IAB.
	//    Additionally, a correct GVL is not malformed - i,e, the entire response content sent from IAB is all there.
	// 2. A correct GVL encoding is one from which the correct response content sent from IAB can be rebuilt
	// For this code to work, we have to assume IAB's server is returning the correct result. From that point forward, we have to be able to cache that response in a form
	// that's retreivable, and if retreived, we can rebuild the original response the encoding retrieved

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Did not succeed in created byte array representation of body.")
		return err
	}
	// Error Case 1: body is nil
	if body == nil {
		// The slice being nil (empty for this type) means that the EOF character was the only character in the body and therefore,
		// the server did not return a response that we could work with. Therefore, there is an issue with the third party libraries used,
		// the link is outdated, or the IAB server is experiencing issues
		return errors.New("Response body is not returned in call")
	}

	// Error Case 2: body is malformed
	// The documentation for ioutil.ReadAll states explicitly states that reaching an EOF is exactly how the byte array return is parsed.
	// Therefore, it could be possible that if part of the file is not sent b/c of a networking issue, then an EOF could be placed
	// in the middle of the resppnse, and not at the end. Therefore, simply the error returning as empty here does not mean that the response
	// content retreived from the IAB server was correct. We have to add a check here to see if the response body was returned as defined within
	// the technical specification.
	if gvl.isIABResponseBodyMalformed(body) == true {
		return errors.New("Server response body is malformed")
	}

	// If body is neither nil nor malformed, then any correctness errors will be in terms of encoding the response content incorrectly as defined by the constraints
	json.Unmarshal(body, &gvl)
	return nil
}

func (gvl *GVLVersionTwoValue) storeGVLVersion2ValueIntoCache(resp *http.Response) bool {
	// Determine the number of seconds to cache the GVL
	expiryTime, err := gvl.getCachingPeriodOfGVLInSeconds(resp)
	if err != nil {
		// There is a serious problem here and we can't cache the response
		log.Print("Wss not able to cache the response.")
		return false
	}

	// use the number of seconds to determine the cache expiry time, and use that to store the cookie
	err = ezcache.ReplaceKeyObject("middleton", gvlCacheKey, gvl, int32(expiryTime))
	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

func (gvl *GVLVersionTwoValue) getCachingPeriodOfGVLInSeconds(resp *http.Response) (int, error) {
	// Determine the number of days to cache the GVL
	maxAge := resp.Header.Get("cache-control")
	if maxAge == "" {
		return 0, errors.New("returned error for cache-control header")
	}
	// the return should be in the format of max-age=<some integer amount>
	maxAgeResult := strings.Split(maxAge, "=")[1]
	maxAgeResultInt, err := strconv.Atoi(maxAgeResult)
	if err != nil {
		// can't set expiry time to an integer
		return 0, errors.New("can't set expiry time to an integer")
	}
	return maxAgeResultInt, nil
}
