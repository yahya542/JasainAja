

package models

type Provider struct {
    Provider_id int `json:"provider_id"`
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
    Skills string `json:"skills"`
    Bio string `json:"bio"`
    Rating float64 `json:"rating"`
}
