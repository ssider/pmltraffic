package body

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	k8sDebugEnvs = []string{
		"POD_NAME",
		"POD_NAMESPACE",
	}
)

func K8sDebugBody() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("time: %s \n", time.Now().Format("2006-01-02 15:04:05.000")))
	buf.WriteString("envs:\n")
	for _, env := range k8sDebugEnvs {
		buf.WriteString(fmt.Sprintf("    %s: %s\n", env, os.Getenv(env)))
	}
	return buf.String()
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	delayStr := r.URL.Query().Get("delay")
	if delay, err := strconv.Atoi(delayStr); err == nil {
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	rspCodeStr := r.URL.Query().Get("rcode")
	if rspCode, err := strconv.Atoi(rspCodeStr); err == nil {
		w.WriteHeader(rspCode)
	}
	fmt.Fprintf(w, "protocol: %s request uri %s  path: %s \n", r.Proto, r.Host, r.RequestURI)
	fmt.Fprint(w, K8sDebugBody())
}
