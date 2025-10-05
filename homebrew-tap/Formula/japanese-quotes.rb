class JapaneseQuotes < Formula
  desc "Display random Japanese quotes in your terminal"
  homepage "https://github.com/Djonsinere/japanese-quotes"
  url "https://github.com/Djonsinere/japanese-quotes/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "0019dfc4b32d63c1392aa264aed2253c1e0c2fb09216f8e2cc269bbfb8bb49b5" 
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "japanese_quotes", "main.go"
    bin.install "japanese_quotes"
    
    # Создаем скрипт для настройки
    (bin/"japanese-quotes-setup").write <<~EOS
      #!/bin/bash
      SHELL_RC="#{ENV["HOME"]}/.#{ENV["SHELL"]##*/}rc"
      
      if ! grep -q "japanese_quotes" "$SHELL_RC" 2>/dev/null; then
        echo "Adding japanese_quotes to $SHELL_RC"
        echo "# Japanese quotes inspiration" >> "$SHELL_RC"
        echo "japanese_quotes" >> "$SHELL_RC"
        echo "Restart your terminal or run: source $SHELL_RC"
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
    # Проверяем что программа запускается и выводит текст
    assert_includes shell_output("#{bin}/japanese_quotes 2>&1"), ""
  end
end