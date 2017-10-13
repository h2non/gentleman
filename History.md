
v2.0.3 / 2017-10-13
===================

  * fix(#37): header copy in redirect plugin

v2.0.2 / 2017-09-20
===================

  * fix(#36): make middleware layer thread-safe
  * feat(docs): add sponsor ad

v2.0.1 / 2017-09-14
===================

  * fix(#35): do not stop on parent middleware layer in error phase
  * fix(travis): gentleman requires Go 1.7+
  * fix(examples): formatting issue
  * fix(examples): formatting issue
  * Merge branch 'master' of https://github.com/h2non/gentleman
  * feat(#33): add error handling example
  * refactor(docs): split community based plugins
  * Merge pull request #34 from izumin5210/logger
  * Add link for gentleman-logger
  * Merge pull request #31 from djui/patch-1
  * Fix typo
  * feat(docs): add gock reference

v2.0.0 / 2017-07-26
===================

  * fix(merge): resolve conflicts from master
  * feat(docs): add versions notes
  * refactor(docs): update version badge
  * refactor(docs): update version badge
  * Merge branch 'master' of https://github.com/h2non/gentleman into v2
  * rerafctor(docs): use v2 godoc links
  * refactor(examples): normalize import statements
  * fix(bodytype): add Type() alias public method
  * feat(History): update changes
  * feat(version): bump to v2.0.0-rc.0
  * fix: format code accordingly
  * fix(travis): use Go 1.7+
  * feat(v2): push v2 preview release

v2.0.0-rc.0 / 2017-03-18
========================

  * feat(version): bump to v2.0.0-rc.0
  * fix: format code accordingly
  * fix(travis): use Go 1.7+
  * refactor(context): adopt standard `context`. Introduces several breaking changes
  * feat(v2): push v2 preview release

v1.0.4 / 2017-05-31
===================

  * Merge pull request #28 from eyalpost/master
  * refactor: remove finalizers (this introduces a minor breaking change)
  * feat(docs): add v2 notice

v1.0.3 / 2017-03-17
===================

  * fix(#23): persist context data across body updates
  * fix(lint): use proper code style
  * fix(test): several lint issues
  * fix(#22): adds support for multipart multiple form values. This introduces a minor interface breaking change
  * feat(travis): add Go 1.8 CI support

v1.0.2 / 2017-02-22
===================

  * Merge pull request #21 from ewilazarus/master
  * fix(#20): Prevents finalizing chunked responses

## 1.0.1 / 2016-01-29

- fix(body): do not enforce POST method unless method is not already present.

## 0.1.3 / 03-03-2016

- fix(#12): define request method via middleware.

## 0.1.2 / 28-02-2016

- feat(body): transparently define POST method if needed.
- feat(multipart): transparently define POST method if needed.

## 0.1.1 / 27-02-2016

- feat(multipart): support custom form field name.

## 0.1.0 / 25-02-2016

- First release.
