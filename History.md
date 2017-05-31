
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
