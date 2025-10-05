class JapaneseQuotes < Formula
  desc "Display random Japanese quotes in your terminal"
  homepage "https://github.com/Djonsinere/japanese-quotes"
  url "https://github.com/Djonsinere/japanese-quotes/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "0019dfc4b32d63c1392aa264aed2253c1e0c2fb09216f8e2cc269bbfb8bb49b5"
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