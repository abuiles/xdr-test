require 'bundler'
Bundler.setup()

namespace :xdr do

  # As stellar-core adds more .x files, we'll need to update this array
  # Prior to launch, we should be separating our .x files into a separate
  # repo, and should be able to improve this integration.
  HAYASHI_XDR = [
                 "foo.x",
                ]

  task :generate do
    require "pathname"
    require "xdrgen"

    paths = Pathname.glob("src/**/*.x")
    compilation = Xdrgen::Compilation.new(
      paths,
      output_dir: "src/generated",
      namespace:  "stellar-xdr",
      language:   :javascript
    )
    compilation.compile
  end
end
