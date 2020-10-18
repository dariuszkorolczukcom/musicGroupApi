package preset

import (
	"context"
	"fmt"

	DB "github.com/dariuszkorolczukcom/musicGroupApi/util/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Preset struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	LowBand      bool               `bson:"low_band," json:"low_band,"`
	LowPeakShelf string             `bson:"low_peak/shelf," json:"low_peak/shelf,"`
	LowFreqHz    string             `bson:"low_freq_hz," json:"low_freq_hz,"`
	LowGain      string             `bson:"low_gain," json:"low_gain,"`
	LowMidBand   bool               `bson:"low_mid_band," json:"low_mid_band,"`
	LowMidHiLowQ string             `bson:"low_mid_hi_low_q," json:"low_mid_hi_low_q,"`
	LowMidFreqHz int                `bson:"low_mid_freq_hz," json:"low_mid_freq_hz,"`
	LowMidGain   int                `bson:"low_mid_gain," json:"low_mid_gain,"`
	HiMidBand    bool               `bson:"hi_mid_band," json:"hi_mid_band,"`
	HiMidFreqKhz float64            `bson:"hi_mid_freq_khz," json:"hi_mid_freq_khz,"`
	HiMidGain    string             `bson:"hi_mid_gain," json:"hi_mid_gain,"`
	HiBand       bool               `bson:"hi_band," json:"hi_band,"`
	HiPeakShelf  string             `bson:"hi_peak_shelf," json:"hi_peak_shelf,"`
	HiFreqKhz    float64            `bson:"hi_freq_khz," json:"hi_freq_khz,"`
	HiGain       int                `bson:"hi_gain," json:"hi_gain,"`
}

func FetchPreset(id string) (Preset, error) {
	var item Preset
	e := DB.Presets.FindOne(context.TODO(), bson.M{"_id": getObjectID(id)}, options.FindOne()).Decode(&item)
	CheckError(e)
	return item, nil
}

func FetchPresets() ([]Preset, error) {
	var items []Preset
	cursor, e := DB.Presets.Find(context.TODO(), bson.D{}, options.Find())
	cursor.All(context.TODO(), &items)
	CheckError(e)
	return items, nil
}

func CreatePreset(p Preset) (*Preset, error) {
	res, e := DB.Presets.InsertOne(context.TODO(), p)
	CheckError(e)
	p.ID = res.InsertedID.(primitive.ObjectID)
	return &p, nil
}

func UpdatePreset(id string, u Preset) (Preset, error) {

	pByte, e := bson.Marshal(u)
	fmt.Println(u)
	CheckError(e)

	var update bson.M
	e = bson.Unmarshal(pByte, &update)
	fmt.Println(update)
	CheckError(e)

	_, e = DB.Presets.UpdateOne(context.TODO(), bson.M{"_id": getObjectID(id)}, bson.D{{Key: "$set", Value: update}})
	CheckError(e)

	return FetchPreset(id)
}

func DeletePreset(id string) error {
	_, e := DB.Presets.DeleteOne(context.TODO(), bson.M{"_id": getObjectID(id)}, options.Delete())
	CheckError(e)
	return nil
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func getObjectID(s string) primitive.ObjectID {
	objID, e := primitive.ObjectIDFromHex(s)
	CheckError(e)
	return objID
}
