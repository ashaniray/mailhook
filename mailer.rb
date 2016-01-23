#!/usr/bin/env ruby

require 'net/smtp'

msgstr = <<END_OF_MESSAGE
From: Your Name <bob@example.com>
To: Destination Address <alice@example.com>
Subject: test message
Date: Sat, 23 Jun 2001 16:26:43 +0900
Message-Id: <unique.message.id.string@example.com>

This is a test message.
END_OF_MESSAGE

require 'net/smtp'
Net::SMTP.start('localhost', 2025) do |smtp|
  smtp.send_message msgstr, 'bob@example.com', 'alice@example.com'
end
