package main

type Client struct{}

func Shadowing(tracing bool) (*Client, error) {
	var cl *Client

	// Here we are shadowing beacuse we are redeclaring the variable cl inside a inner scope
	if tracing {
		cl, err := clientWithTracing()
		if err != nil {
			return nil, err
		}

		_ = cl // Here we need to reassign this variable, so Go does not throw an error (This makes clear that we are using an inner scope variable)
	} else {
		cl, err := clientWithoutTracing()
		if err != nil {
			return nil, err
		}

		_ = cl
	}

	return cl, nil
}

func ShadowingPossibleSolution(tracing bool) (*Client, error) {
	var (
		cl *Client
		err error
	)

	if tracing {
		if cl, err = clientWithTracing(); err != nil {
			return nil, err
		}
	} else {
		if cl, err = clientWithoutTracing(); err != nil {
			return nil, err
		}
	}

	return cl, err
}

func ShadowingReassignmentSolution(tracing bool) (*Client, error) {
	var client *Client

	if tracing {
		cl, err := clientWithTracing()
		if err != nil {
			return nil, err
		}

		client = cl
	} else {
		cl, err := clientWithoutTracing()
		if err != nil {
			return nil, err
		}

		client = cl
	}

	return client, nil
}

func ShadowingPossibleSolutionOther(tracing bool) (*Client, error) {
	var (
		cl *Client
		err error
	)

	if tracing {
		cl, err = clientWithTracing()
	} else {
		cl, err = clientWithoutTracing()
	}

	return cl, err
}

func ShadowingPossibleSolutionOther2(tracing bool) (*Client, error) {
	if tracing {
		return clientWithTracing()
	} else {
		return clientWithoutTracing()
	}
}

func clientWithTracing() (*Client, error) {
	return &Client{}, nil
}

func clientWithoutTracing() (*Client, error) {
	return &Client{}, nil
}
