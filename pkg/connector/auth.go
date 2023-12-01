package connector

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type AuthenticationMethod func(tenantId, clientId string) (azcore.TokenCredential, error)

func WithClientSecret(clientSecret string) AuthenticationMethod {
	return func(tenantId, clientId string) (azcore.TokenCredential, error) {
		return azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
	}
}

func WithClientCertificate(clientCertificatePath string) AuthenticationMethod {
	return func(tenantId, clientId string) (azcore.TokenCredential, error) {
		file, err := os.Open(clientCertificatePath)
		if err != nil {
			return nil, wrapError(err, "error opening client certificate file")
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			return nil, wrapError(err, "error getting client certificate file info")
		}

		bytes := make([]byte, fileInfo.Size())
		_, err = file.Read(bytes)
		if err != nil {
			return nil, wrapError(err, "error reading client certificate file")
		}

		certs, key, err := azidentity.ParseCertificates(bytes, nil)
		if err != nil {
			return nil, wrapError(err, "error parsing client certificate")
		}

		return azidentity.NewClientCertificateCredential(tenantId, clientId, certs, key, nil)
	}
}
