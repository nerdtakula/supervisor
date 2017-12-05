package supervisor

// Read length from the stdout for specified process
func (c Client) ReadProcessStdout(name string, offset, length int) (string, error) {
  var result string
  err := c.makeRequest("supervisor.readProcessStdoutLog", []interface{}{name, offset, length}, &result)
  if err != nil {
    return "", err
  }
  return result, nil
}
