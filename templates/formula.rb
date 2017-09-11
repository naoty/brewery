class {{.ClassName}} < Formula
  desc     "{{.Desc}}"
  homepage "{{.Homepage}}"
  url      "{{.URL}}"
  sha256   "{{.Sha256}}"

  def install
    bin.install "{{.Name}}"
  end
end
