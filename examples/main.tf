terraform {
  required_providers {
    playgroundtech = {
      version = "0.2"
      source   = "playgroundtech.io/job/application"
    }
  }
}


provider "playgroundtech" {
  email    = "email@email.com"
  password = "Hejsan123Hejsan123"
}

resource "playgroundtech_application" "test" {
  phone_number = "0123456789"
  email        = "email@email.com"
  linkedin     = "linkedin.com"
  github       = "github.com"
  homepage     = "homepage.com"
}
