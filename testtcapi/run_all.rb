Dir[File.dirname(File.absolute_path(__FILE__)) + '/tests/test_*.rb'].each {|file| require file }

STDOUT.sync = true

puts 'run_all.rb - done'
