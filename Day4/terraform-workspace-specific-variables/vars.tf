variable "environment" {
  type = string
  default = "default" 
}

variable "message" {
   description = "Variable that can hold workspace specific value"
   type = string
   default = "dev"
}

variable "common_variable" {
    description = "Common variable accessible to all workspaces"
    type = string
    default = "This is a common string variable"
}
