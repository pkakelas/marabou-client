package main

func HelloController(req HelloMsg) (error, map[string]interface{}) {
	if req.Version == "" || len(req.Version) < 3 {
		return InvalidInputError, nil
	}
	if req.Version[0:3] != "0.1" {
		return IncopatibleVersions, nil
	}

	res := map[string]interface{}{
		"version": "0.1.2",
		"agent":   "Mitsos-Core Client 0.7",
	}

	return nil, res
}
