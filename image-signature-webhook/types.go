package main

// v1Beta1ListOccurrencesResponse is the JSON response from occurrencesPath
type v1Beta1ListOccurrencesResponse struct {
	Occurrences   []v1Beta1Occurrence `json:"occurrences"`
	NextPageToken string              `json:"nextPageToken"`
}

// v1Beta1Occurrence is the JSON structure of an occurrence
type v1Beta1Occurrence struct {
	Name     string `json:"name"`
	Resource struct {
		Name        string `json:"name"`
		URI         string `json:"uri"`
		ContentHash string `json:"contentHash"`
	} `json:"resource"`
	NoteName    string `json:"noteName"`
	Kind        string `json:"kind"`
	Remediation string `json:"remediation"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	Attestation struct {
		Attestation struct {
			PgpSignedAttestation struct {
				Signature   string `json:"signature"`
				ContentType string `json:"contentType"`
				PgpKeyID    string `json:"pgpKeyId"`
			}
		} `json:"attestation"`
	} `json:"attestation"`
}
