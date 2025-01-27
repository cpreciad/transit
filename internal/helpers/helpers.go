package helpers

import(
    "time"
)

func UTCtoPST(utcTime string) (time.Time, error){

    pst, err := time.LoadLocation("America/Los_Angeles")
    if err != nil{
        return time.Time{}, err
    }

    t, err := time.Parse(time.RFC3339, utcTime)
    if err != nil{
        return time.Time{}, err
    }
    t = t.In(pst)

    return t, nil
}
