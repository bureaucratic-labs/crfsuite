#include "crfsuite/include/crfsuite.h"

crfsuite_model_t* NewModelFromFile(char* path) {
    crfsuite_model_t* model = NULL;
    int ret = crfsuite_create_instance_from_file(path, (void**)&model);
    return model;
}

crfsuite_dictionary_t* GetModelLabels(crfsuite_model_t* model) {
    crfsuite_dictionary_t* labels = NULL;
    int ret = model->get_labels(model, &labels);
    return labels;
}


crfsuite_dictionary_t* GetModelAttributes(crfsuite_model_t* model) {
    crfsuite_dictionary_t* attributes = NULL;
    int ret = model->get_attrs(model, &attributes);
    return attributes;
}

int DictionaryLength(crfsuite_dictionary_t* dictionary) {
    return dictionary->num(dictionary);
}

int DictionaryGet(crfsuite_dictionary_t* dictionary, const char *str) {
    return dictionary->get(dictionary, str);
}

int DictionaryToID(crfsuite_dictionary_t* dictionary, const char *str) {
    return dictionary->to_id(dictionary, str);
}

