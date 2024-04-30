package getdayfunc

import "errors"

func GetDay(d int) (string, error) {

	if d < 1 || d > 7 {
		return "", errors.New("invalid number of day")
	} else {
		switch d {
		case 1:
			return "Day is Monday", nil
		case 2:
			return "Day is Tuesday", nil
		case 3:
			return "Day is Wednesday", nil
		case 4:
			return "Day is Thursday", nil
		case 5:
			return "Day is Friday", nil
		case 6:
			return "Day is Saturday", nil
		case 7:
			return "day is Sunday", nil
		}
	}
	return "", nil
}
