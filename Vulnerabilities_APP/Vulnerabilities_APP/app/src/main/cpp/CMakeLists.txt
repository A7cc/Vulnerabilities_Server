# For more information about using CMake with Android Studio, read the
# documentation: https://d.android.com/studio/projects/add-native-code.html.
# For more examples on how to use CMake, see https://github.com/android/ndk-samples.

cmake_minimum_required(VERSION 3.22.1)

project(vulnerabilities)

add_library(vuln_app SHARED vuln_app.c encrypt.c)
add_library(fridastudy SHARED fridastudy.c)
find_library(log-lib log)

target_link_libraries(vuln_app ${log-lib})
target_link_libraries(fridastudy ${log-lib} dl)