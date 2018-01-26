package supervisor

// Get info about a specific process
func (c Client) GetProcessInfo(name string) (*ProcessInfo, error) {
	result := &ProcessInfo{}
	err := c.makeRequest("supervisor.getProcessInfo", []interface{}{name}, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get info about all processes
func (c Client) GetAllProcessInfo() ([]ProcessInfo, error) {
	var result []ProcessInfo
	err := c.makeRequest("supervisor.getAllProcessInfo", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Start a process
func (c Client) StartProcess(name string, wait bool) (bool, error) {
	var result bool
	err := c.makeRequest("supervisor.startProcess", []interface{}{name, wait}, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Start all processes listed in the group specified
func (c Client) StartProcessGroup(name string, wait bool) ([]ProcessInfo, error) {
	var result []ProcessInfo
	err := c.makeRequest("supervisor.startProcessGroup", []interface{}{name, wait}, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Start all processes listed in the configuration file
func (c Client) StartAllProcesses(wait bool) ([]ProcessInfo, error) {
	var result []ProcessInfo
	err := c.makeRequest("supervisor.startAllProcesses", wait, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Stop a process
func (c Client) StopProcess(name string, wait bool) (bool, error) {
	var result bool
	err := c.makeRequest("supervisor.stopProcess", []interface{}{name, wait}, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Stop all processes listed in the group specified
func (c Client) StopProcessGroup(name string, wait bool) ([]ProcessInfo, error) {
	var result []ProcessInfo
	err := c.makeRequest("supervisor.stopProcessGroup", []interface{}{name, wait}, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Stop all processes listed in the configuration file
func (c Client) StopAllProcesses(wait bool) ([]ProcessInfo, error) {
	var result []ProcessInfo
	err := c.makeRequest("supervisor.stopAllProcesses", wait, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Send an arbitrary UNIX signal to the process specified
func (c Client) SignalProcess(name string, signal int) ([]*ProcessInfo, error) { return nil, nil }

// Send an arbitrary UNIX signal to all process in the group specified
func (c Client) SignalProcessGroup(name string, signal int) ([]*ProcessInfo, error) { return nil, nil }

// Send an arbitrary UNIX signal to all processes listed in the configuration file
func (c Client) SignalAllProcesses(signal int) ([]*ProcessInfo, error) { return nil, nil }

// Send a string of characters to the STDIN of the specified process
func (c Client) SendProcessStdin(name, chars string) error { return nil }

// Send an event that will be received by any event listener subprocesses
// subscribing to the 'RemoteCommunicationEvent'
func (c Client) SendRemoteCommEvent(t, data string) error { return nil }

// Reload the configuration
func (c Client) Reload() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := c.makeRequest("supervisor.reloadConfig", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update the config for a running process from the configuration file
func (c Client) AddProcessGroup(name string) (bool, error) {
	var result bool
	err := c.makeRequest("supervisor.addProcessGroup", name, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Remove a stopped process group from the active configuration
func (c Client) RemoveProcessGroup(name string) (bool, error) {
	var result bool
	err := c.makeRequest("supervisor.removeProcessGroup", name, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
