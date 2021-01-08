#
# Cookbook:: demo_book
# Recipe:: default
#
# Copyright:: 2019, The Authors, All Rights Reserved.

file "D:/NAVYA/DevOps/Interns2019/source/chef_file1.txt" do 
    content 'hello this is the first chef file'
    action :create_if_missing
    
end