## Image updater for Github Actions
This is a utility tool written in Go, designed to update a specific attribute in a Helm chart YAML file. Specifically, it modifies the image tag attribute. The Helm chart path and the image tag to be used are provided through environment variables.


## Usage
We support two methods for updating image tags:

If your Helm chart is located in the same repository as your code, you can incorporate the following step into your workflow:

```yaml

    - name: Update Image tag 
      id: image-updater
      uses: pooriaghaedi/image-updater-action@main
      env:
        IMAGE_TAG: "v1.0.2"
        CHART_PATH: yourchart/values.yaml
```

If your Helm chart is located in a different repository, you'll need to clone that repository as a part of your workflow.


```yaml
    - name: Checkout private tools
      uses: actions/checkout@v3
      with:
        repository: youruser/repositoryname
        token: ${{ secrets.GH_PAT }} # `GH_PAT` is a secret that contains your PAT
        path: charts # Directory into which the repository should be cloned.
        ref: main # preffered branch name

    - name: Update Image tag 
      id: image-updater
      uses: pooriaghaedi/image-updater-action@main
      env:
        IMAGE_TAG: "v1.0.2"
        CHART_PATH: charts/yourchart/values.yaml
```

### Tip
You can generate image tags based on commit hash:

```yaml
    - name: Generate image tag
      id: tags
      run: |
        echo "::set-output name=sha_short::$(git rev-parse --short HEAD)"

    - name: Update Image tag 
      id: image-updater
      uses: pooriaghaedi/image-updater-action@main
      env:
        IMAGE_TAG: ${{ steps.tags.outputs.sha_short }}
        CHART_PATH: yourchart/values.yaml
```


## Contributing
If you wish to contribute to this project by adding more features or improving the existing ones, please feel free to open a pull request.

## License
This project is released under the MIT license.