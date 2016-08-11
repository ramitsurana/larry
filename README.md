# Larry

Move Infrastructure easily between Clouds using Terraform

![larry](https://cloud.githubusercontent.com/assets/8342133/17511593/39d516c6-5e41-11e6-8fda-a57850333d9d.png)

Truck Logo &copy; [Igor Gamajunov](https://dribbble.com/cybertigro)

## What is Larry ?

Larry is a tool to help you shift infrastructure across different clouds using Terraform

## Features

* Supports Multiple Cloud Providers
* Uses [Terraform](http://terraform.io) at the backend for provisioning


### Requirements

* Go v1.6

(Note that the code uses the relatively new Go vendoring, so building requires Go 1.6 or later, or you must export GO15VENDOREXPERIMENT=1 when building with Go 1.5.) 

### How it works

Basically,Larry uses Terraform for its provisioning and working.There are seprate commands using which we can create and destroy  works in 5 simple steps:

**Step 1** Larry fetch

Using this command Larry will automatically fetch all the information from your cloud provider.

**Step 2** View all the information

Using the created json file you can view all the information of your current infrastructure and edit it if you want.

**Step3** Larry convert

Larry converts the json file configurations to the shifting cloud provider

**Step 4** Larry deploy

This cmd will create the infrastructure as mentioned in the converted json file 

**Step 5** Larry check

Remember to check your infrastructure status of the cloud to which you have shifted your setup

### Getting Started

Releasing Soon

## Contributing

Contributions can be made easily by making PR's and opening issues on the github repo.
Big Thank you to all the [contributors][3] for your awesome contributions !

## License

[MIT License](LICENSE)

[1]: http://docker.com
[2]: http://ramitsurana.github.io/turbo
[3]: https://github.com/ramitsurana/larry/graphs/contributors
