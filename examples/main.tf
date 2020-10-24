terraform {
  required_providers {
    playgroundtech = {
      versions = ["0.2"]
      source   = "playgroundtech.se/edu/playgroundtech"
    }
  }
}


provider "playgroundtech" {
  email    = "email@email.com"
  password = "password"
}

resource "playgroundtech_application" "test" {
  phone_number = "0123456789"
  email        = "email@email.com"
  linkedin     = "linkedin.com"
  github       = "github.com"
  homepage     = "homepage.com"
}
