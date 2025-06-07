package pkg

import (
    ffmappkg "ffmap/pkg"
    chartspkg "charts/pkg"
)

func Use() error {
    _, err := ffmappkg.Open("test.ffmap")
    if err != nil {
        return err
    }
    _, _, _, err = chartspkg.PopulateData()
    return err
}
