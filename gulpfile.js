const path = require('path')
const gulp = require('gulp')
const util = require('gulp-util')
const notifier = require('node-notifier')
const sync = require('gulp-sync')(gulp).sync
const child = require('child_process')
const os = require('os')

const platform = os.platform()
let server = null
let isServerBuildError = false
let pathFolder
if (platform === 'win32') {
  // Windows
  pathFolder = __dirname.split('\\')
} else {
  // Linux / MacOS
  pathFolder = __dirname.split('/')
}
// Application name
const app = pathFolder[pathFolder.length - 1]

const mainPackageDir = './cmd/' + app
const publicDir = 'public'

// Compile application
gulp.task('server:build', () => {
  // Build application in the "gobin" folder
  let build = child.spawnSync('go', ['install', mainPackageDir])
  // let build = child.spawnSync('go', ['build', mainPackageDir])
  isServerBuildError = false
  if (build.stderr.length) {
    util.log(util.colors.red('Something wrong with this version :'))
    let lines = build.stderr.toString()
      .split('\n').filter((line) => {
        return line.length
      })

    for (let l in lines) {
      util.log(util.colors.red(
        'Error (go install): ' + lines[l]
      ))
    }

    isServerBuildError = true

    notifier.notify({
      title: 'Error (go install)',
      message: lines.toString(),
      icon: path.join(__dirname, '../misc/notifier/gopher.png'), // Absolute path (doesn't work on balloons)
      sound: true, // Only Notification Center or Windows Toasters
    })
  }
  return build
})

// Server launch
gulp.task('server:spawn', () => {
  // Stop the server
  if (server && server !== null) {
    server.kill()
  }

  // Run the server
  if (!isServerBuildError) {
    if (platform === 'win32') {
      server = child.spawn(app + '.exe')
    } else {
      server = child.spawn(app)
    }

    server.on('error', err => util.log(util.colors.red(err)))

    // Display terminal informations
    server.stderr.on('data', data => process.stdout.write(data.toString()))
    server.stdout.on('data', data => process.stdout.write(data.toString()))
  }
})

// Watch files
gulp.task('server:watch', ['server:build', 'server:spawn'], () => {
  gulp.watch([
    '*.go',
    '**/*.go',
  ], sync([
    'server:build',
    'server:spawn'
  ]))
  // gulp.watch(publicDir + '/**/*.html', ['html'])
})

gulp.task('html', () => {
  gulp.src(publicDir + '/**/*.html')
})

gulp.task('default', ['server:watch'])
