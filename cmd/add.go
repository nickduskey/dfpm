package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/manifoldco/promptui"

	"github.com/nickduskey/dfpm/services"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

type DependencyTree struct {
	Dependencies []string
}

func (d *DependencyTree) HasDependency(packageName string) bool {
	for _, dep := range d.Dependencies {
		if dep == packageName {
			return true
		}
	}
	return false
}

func (d *DependencyTree) FindNewDependencies(dependencyTree map[string]string) []string {
	var newDependencies []string
	for key, _ := range dependencyTree {
		if !d.HasDependency(key) {
			newDependencies = append(newDependencies, key)
		}
	}
	return newDependencies
}

func (d *DependencyTree) GetTotalDependencies(packageName string) (int, error) {
	s := spinner.New(spinner.CharSets[12], 100*time.Millisecond)
	s.Suffix = " Fetching all the dependencies. This may take a while..."
	s.Color("blue", "bold")
	s.FinalMSG = "Whew! Done finding all the dependencies!\n"
	s.Start()
	npmsService := services.Npms{
		Version: "0.1.0",
	}

	res, err := npmsService.GetPackageDetails(packageName)
	if err != nil {
		return 0, err
	}

	rootDeps := []string{}
	for key := range res.Collected.Metadata.Dependencies {
		// TODO check if we already have this dep
		res, err := npmsService.GetPackageDetails(key)
		if err != nil {
			return 0, err
		}
		rootDeps = append(rootDeps, d.FindNewDependencies(res.Collected.Metadata.Dependencies)...)
	}
	newDepsThisRound := len(rootDeps)
	depsToCheck := []string{}
	depsToCheck = append(depsToCheck, rootDeps...)
	for {
		for _, dep := range depsToCheck {
			res, err := npmsService.GetPackageDetails(dep)
			if err != nil {
				return 0, err
			}
			newDeps := d.FindNewDependencies(res.Collected.Metadata.Dependencies)
			d.Dependencies = append(d.Dependencies, newDeps...)
			newDepsThisRound = len(newDeps)
			depsToCheck = append(depsToCheck, newDeps...)
		}
		if newDepsThisRound == 0 {
			break
		}
	}
	s.Stop()
	return len(d.Dependencies), nil
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Carefully add a package to your javascript project",
	Long:  "Carefully add a package to your javascript project, fully aware of the costs. There is no such thing as free lunch",
	Run: func(_ *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Printf("Got arg: %s\n", args[0])

			var tree DependencyTree

			totalDeps, err := tree.GetTotalDependencies(args[0])
			if err != nil {
				log.Fatalln(err)
			}
			q := fmt.Sprintf("About to install %d dependencies, do you really want to do that???", totalDeps)

			prompt := promptui.Select{
				Label: q,
				Items: []string{"Yes", "No"},
			}

			_, result, err := prompt.Run()
			if err != nil {
				log.Fatalf("Failed getting prompt back from user %v\n", err)
			}
			if result == "Yes" {
				registryService := services.Registry{}
				tarball, err := registryService.GetPackageTarballURL(args[0])
			}
			log.Println("Yeah I probably wouldn't install that dependency either.")
			return
		} else {
			fmt.Println("Required arg not passed in")
		}
	},
}
