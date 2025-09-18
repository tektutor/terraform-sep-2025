package provider

import (
   "context"
   "log"
   "os"
   "io"

   "github.com/moby/moby/client"

   "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
   "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDockerImage() *schema.Resource {
	return &schema.Resource {

		Description: "This resource will support Create, Read, Update and Delete docker image resources via Terraform.",

		CreateContext: resourceCreateDockerImage,
		ReadContext:   resourceReadDockerImage,
		UpdateContext: resourceUpdateDockerImage,
		DeleteContext: resourceDeleteDockerImage,

		Schema: map[string]*schema.Schema {

			"image_name": {
				Description: "Name of the docker image including tag",
				Type       : schema.TypeString,
				Required   : true,
			},
		},
	}
}

func resourceCreateDockerImage( ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	imageName := d.Get("image_name").(string)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	reader, err1 := cli.ImagePull(ctx,imageName, client.ImagePullOptions{})
	if err1 != nil {
		panic(err1)
	}

	io.Copy(os.Stdout, reader)

	resp, err2 := cli.ImageInspect(ctx, imageName)
	if err2 != nil {
		panic(err2)
	}

	d.SetId(resp.ID)

	return nil
}


func resourceReadDockerImage( ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func resourceUpdateDockerImage( ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func resourceDeleteDockerImage( ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	imageName := d.Get("image_name").(string)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	_, err = cli.ImageRemove( ctx, imageName, client.ImageRemoveOptions{Force: true} )
	if err != nil {
		log.Printf("Unable to delete image : %s", err)
		panic(err)
	}


	return nil
}
