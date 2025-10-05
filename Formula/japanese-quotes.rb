class JapaneseQuotes < Formula
  desc "Display random Japanese quotes in your terminal"
  homepage "https://github.com/Djonsinere/homebrew-jquotes.git"
  url "https://github.com/yourusername/japanese-quotes/archive/v1.0.0.tar.gz"
  sha256 "PUT_SHA256_HERE" 
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "japanese_quotes"
    bin.install "japanese_quotes"
    
    (bin/"japanese-quotes-setup").write <<~EOS
      #!/bin/bash
      SHELL_RC="#{ENV["HOME"]}/.#{ENV["SHELL"]##*/}rc"
      
      if ! grep -q "japanese_quotes" "$SHELL_RC" 2>/dev/null; then
        echo "Adding japanese_quotes to $SHELL_RC"
        echo "# Japanese quotes inspiration" >> "$SHELL_RC"
        echo "japanese_quotes" >> "$SHELL_RC"
      else
        echo "japanese_quotes already in shell config"
      fi
    EOS
    
    chmod 0755, bin/"japanese-quotes-setup"
  end

  def caveats
    <<~EOS
      To automatically run japanese-quotes when you open your terminal, run:
        japanese-quotes-setup
      
      Or manually add this line to your ~/.zshrc or ~/.bashrc:
        japanese_quotes
    EOS
  end

  test do
    assert_match "七転び八起き", shell_output("#{bin}/japanese_quotes 2>&1", 1)
  end
end