class JapaneseQuotes < Formula
  desc "Display random Japanese quotes in your terminal"
  homepage "https://github.com/Djonsinere/japanese-quotes"
  url "https://github.com/Djonsinere/homebrew-jquotes/archive/refs/tags/v1.0.1.tar.gz"
  sha256 "ce2132f1fc2937390249500bd9678d04d30a0a127bf84a4761617cc5e9a9b09a"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "japanese_quotes"
    bin.install "japanese_quotes"
  end

  def caveats
    <<~EOS
      To show Japanese quotes when opening your terminal, add to your ~/.zshrc :
        japanese_quotes
    EOS
  end

  test do
    system "#{bin}/japanese_quotes"
  end
end